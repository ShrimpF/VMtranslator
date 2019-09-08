package codeWriter

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ShrimpF/vmTranslator/enum"
)

// arithmeticID -- indentify each jpm-label
var arithmeticID int

// callID -- identify each call
var callID int

// CodeWriter -
type CodeWriter struct {
	fileName string
	writer   *os.File
}

// New -- create CodeWriter
func New(name string, w *os.File) *CodeWriter {
	return &CodeWriter{fileName: name, writer: w}
}

func (cw *CodeWriter) Write(code string) {
	fmt.Fprint(cw.writer, code+"\n")
}

// WriteArithmetic -- write results of arithmetic operation
func (cw *CodeWriter) WriteArithmetic(op enum.Operation) {
	switch op {
	case enum.Add:
		cw.Write(arithmeticTpl1())
		cw.Write("M=M+D")
	case enum.Sub:
		cw.Write(arithmeticTpl1())
		cw.Write("M=M-D")
	case enum.And:
		cw.Write(arithmeticTpl1())
		cw.Write("M=M&D")
	case enum.Or:
		cw.Write(arithmeticTpl1())
		cw.Write("M=M|D")
	case enum.Eq:
		cw.Write(arithmeticTpl2("JEQ"))
	case enum.Gt:
		cw.Write(arithmeticTpl2("JGT"))
	case enum.Lt:
		cw.Write(arithmeticTpl2("JLT"))
	case enum.Neg:
		cw.Write(arithmeticTpl3())
		cw.Write("M=-M")
	case enum.Not:
		cw.Write(arithmeticTpl3())
		cw.Write("M=!M")
	}
}

//WritePushPop -- write push and pop assemble code
func (cw *CodeWriter) WritePushPop(cmdType enum.Command, segment string, index int) {
	if cmdType == enum.CPush {
		cw.Write(pushTpl(segment, index, cw.fileName))
	} else if cmdType == enum.CPop {
		cw.Write(popTpl(segment, index, cw.fileName))
	}
}

// WriteLabel -- write label cmd's assebly code
func (cw *CodeWriter) WriteLabel(label string) {
	cw.Write("(" + label + ")")
}

// WriteGoto -- write go-to cmd's assebly code
func (cw *CodeWriter) WriteGoto(label string) {
	cw.Write("@" + label)
	cw.Write("0;JMP")
}

// WriteIf -- write go-if cmd's assembly code
func (cw *CodeWriter) WriteIf(label string) {
	cw.Write("@SP")
	cw.Write("AM=M-1")
	cw.Write("D=M")
	cw.Write("@" + label)
	cw.Write("D;JNE")
}

// WriteCall -- write call assembly code
func (cw *CodeWriter) WriteCall(funcName string, numOfArgs int) {
	ret := funcName + "RET" + strconv.Itoa(callID)
	callID++

	// push return address
	cw.Write("@" + ret)
	cw.Write("D=A")
	cw.Write(pushDtoStackTpl())

	// push @LCL,@ARG,@THIS,@THAT
	for _, address := range []string{"@LCL", "@ARG", "@THIS", "@THAT"} {
		cw.Write(address)
		cw.Write("D=M")
		cw.Write(pushDtoStackTpl())
	}

	// LCL = SP
	cw.Write("@SP")
	cw.Write("D=M")
	cw.Write("@LCL")
	cw.Write("M=D")

	// ARG = SP - n -5
	cw.Write("@" + strconv.Itoa(numOfArgs+5))
	cw.Write("D=D-A")
	cw.Write("@ARG")
	cw.Write("M=D")

	// goto func
	cw.Write("@" + funcName)
	cw.Write("0;JMP")

	//(return address)
	cw.WriteLabel(ret)
}

// WriteFunc -- write func assembly code
func (cw *CodeWriter) WriteFunc(funcName string, numOflocals int) {
	// write label
	cw.WriteLabel(funcName)

	// initialize local variables
	// push 0 , k times
	for i := 0; i < numOflocals; i++ {
		cw.Write("D=0")
		cw.Write(pushDtoStackTpl())
	}

}

// WriteReturn -- write return assembly code
func (cw *CodeWriter) WriteReturn() {
	// frame = LCL
	cw.Write("@LCL")
	cw.Write("D=M")
	cw.Write("@FRAME")
	cw.Write("M=D")

	// ret = *(frame-5)
	cw.Write("@5")
	cw.Write("A=D-A")
	cw.Write("D=M")
	cw.Write("@RET")
	cw.Write("M=D")

	// *ARG=pop()
	cw.Write(popStackToDTpl())
	cw.Write("@ARG")
	cw.Write("A=M")
	cw.Write("M=D")

	// SP = ARG + 1
	cw.Write("@ARG")
	cw.Write("D=M+1")
	cw.Write("@SP")
	cw.Write("M=D")

	// THAT = *(frame-1)
	// THIS = *(frame-2)
	// ARG = *(frame-3)
	// LCL = *(frame-4)
	for i, address := range []string{"@THAT", "@THIS", "@ARG", "@LCL"} {
		cw.Write("@FRAME")
		cw.Write("D=M")
		cw.Write("@" + strconv.Itoa(i+1))
		cw.Write("D=D-A")
		cw.Write("A=D")
		cw.Write("D=M")
		cw.Write(address)
		cw.Write("M=D")
	}

	// goto ret
	cw.Write("@RET")
	cw.Write("A=M")
	cw.Write("0;JMP")
}
