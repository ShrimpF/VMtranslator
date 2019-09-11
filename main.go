package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ShrimpF/vmTranslator/codeWriter"
)

func main() {
	filePath := os.Args[1]

	// parse the first arg
	// if args[1] is .vm-file => convert the .vm-file to .asm-file
	// if args[1] is directory => find all .vm-files and make a array of .vm-file

	if filepath.Ext(filePath) == ".vm" {
		outFileName := strings.Replace(filepath.Base(filePath), ".vm", ".asm", 1)

		outFile, err := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		checkErr(err)
		defer outFile.Close()

		cw := codeWriter.New(filePath, outFile)
		cw.Translate()

	} else {
		outFileName := filepath.Base(filePath) + ".asm"

		outFile, err := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		checkErr(err)
		defer outFile.Close()

		cw := codeWriter.New("", outFile)
		cw.Init()

		for _, path := range makeVMfileArrayFromDir(filePath) {
			fmt.Println("Converting", filepath.Base(path), "to .asm file......")
			cw.SetFileNameAndPath(path)
			cw.Translate()
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
