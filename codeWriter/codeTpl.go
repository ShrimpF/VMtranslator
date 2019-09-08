package codeWriter

import "strconv"

func incrementStackTpl() string {
	return "@SP\n" +
		"M=M+1"
}

func decrementStackTpl() string {
	return "@SP\n" +
		"M=M-1"
}

func pushDtoStackTpl() string {
	return "@SP\n" +
		"A=M\n" +
		"M=D\n" +
		incrementStackTpl()
}

func popStackToDTpl() string {
	return decrementStackTpl() + "\n" +
		"A=M\n" +
		"D=M"
}

// arithmeticTpl1 is template for add/sub/and/or
func arithmeticTpl1() string {
	return "@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1"
}

// arithmeticTpl2 is template for eq/gt/lt
func arithmeticTpl2(jumpType string) string {
	arithmeticID++
	return "@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"@TRUE" + strconv.Itoa(arithmeticID) + "\n" +
		"D;" + jumpType + "\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=0\n" +
		"@COUTINUE" + strconv.Itoa(arithmeticID) + "\n" +
		"0;JMP\n" +
		"(TRUE" + strconv.Itoa(arithmeticID) + ")\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=-1\n" +
		"(COUTINUE" + strconv.Itoa(arithmeticID) + ")"
}

// arithmeticTpl3 is template for not and neg
func arithmeticTpl3() string {
	return "@SP\n" +
		"A=M-1"
}

// pushTpl return push assembler-code
func pushTpl(segment string, index int, fileName string) string {
	segment = convSegmentToSymbol(segment)

	if segment == "constant" {
		return "@" + strconv.Itoa(index) + "\n" +
			"D=A\n" +
			pushDtoStackTpl()
	} else if segment == "temp" {
		return "@" + strconv.Itoa(index+5) + "\n" +
			"D=M\n" +
			pushDtoStackTpl()
	} else if segment == "pointer" {
		return "@" + strconv.Itoa(index+3) + "\n" +
			"D=M\n" +
			pushDtoStackTpl()
	} else if segment == "static" {
		return "@" + fileName + "." + strconv.Itoa(index) + "\n" +
			"D=M\n" +
			pushDtoStackTpl()
	}

	return "@" + strconv.Itoa(index) + "\n" +
		"D=A\n" +
		"@" + segment + "\n" +
		"A=M+D\n" +
		"D=M\n" +
		pushDtoStackTpl()
}

// popTpl return pop assemlber code
func popTpl(segment string, index int, fileName string) string {
	segment = convSegmentToSymbol(segment)

	if segment == "temp" {
		return "@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@" + strconv.Itoa(index+5) + "\n" +
			"M=D"
	} else if segment == "pointer" {
		return "@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@" + strconv.Itoa(index+3) + "\n" +
			"M=D"
	} else if segment == "static" {
		return "@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@" + fileName + "." + strconv.Itoa(index) + "\n" +
			"M=D"
	}

	return "@" + strconv.Itoa(index) + "\n" +
		"D=A\n" +
		"@" + segment + "\n" +
		"D=M+D\n" +
		"@R13\n" +
		"M=D\n" +
		"@SP\n" +
		"AM=M-1\n" +
		"D=M\n" +
		"@R13\n" +
		"A=M\n" +
		"M=D"
}
