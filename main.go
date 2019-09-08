package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ShrimpF/vmTranslator/codeWriter"
	"github.com/ShrimpF/vmTranslator/enum"
	"github.com/ShrimpF/vmTranslator/parser"
)

func main() {
	filePath := os.Args[1]

	// parse the first arg
	// if args[1] is .vm-file => convert the .vm-file to .asm-file
	// if args[1] is directory => find all .vm-files and make a array of .vm-file
	// make directoryname + .asm file to write in asmcode

	if filepath.Ext(filePath) == ".vm" {
		// create a asm-file to write in
		outFileName := strings.Replace(filepath.Base(filePath), ".vm", ".asm", 1)
		outFile, err := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		checkErr(err)
		defer outFile.Close()

		translate(filePath, outFile)
	} else {
		outFileName := filepath.Base(filePath) + ".asm"
		_, err := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		checkErr(err)

		for _, v := range makeVMfileArrayFromDir(filePath) {
			fmt.Println(v)
		}
	}

}

func translate(path string, outputFile *os.File) {
	readFile, err := os.Open(path)
	checkErr(err)
	defer readFile.Close()

	cw := codeWriter.New(filepath.Base(path), outputFile)
	p := parser.New()

	scan := bufio.NewScanner(readFile)
	for scan.Scan() {
		text := removeComment(scan.Text())
		if text == "" {
			continue
		}
		p.SetCode(text)
		switch p.CommandType() {
		case enum.CPush:
			cw.WritePushPop(p.CommandType(), p.Arg1(), p.Arg2())
		case enum.CPop:
			cw.WritePushPop(p.CommandType(), p.Arg1(), p.Arg2())
		case enum.CArithmetic:
			cw.WriteArithmetic(p.Operation())
		case enum.CLabel:
			cw.WriteLabel(p.Arg1())
		case enum.CGoto:
			cw.WriteGoto(p.Arg1())
		case enum.CIf:
			cw.WriteIf(p.Arg1())
		case enum.CFunction:
			cw.WriteFunc(p.Arg1(), p.Arg2())
		case enum.CCall:
			cw.WriteCall(p.Arg1(), p.Arg2())
		case enum.CReturn:
			cw.WriteReturn()
		}
	}
}

func makeVMfileArrayFromDir(path string) []string {
	dir := filepath.Dir(path)
	var vmfiles []string
	files, err := ioutil.ReadDir(path)
	checkErr(err)

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".vm" {
			vmfiles = append(vmfiles, dir+"/"+f.Name())
		}
	}

	return vmfiles
}

func removeComment(text string) string {
	index := strings.Index(text, "//")
	if index != -1 {
		text = text[:index]
	}
	return strings.TrimSpace(text)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
