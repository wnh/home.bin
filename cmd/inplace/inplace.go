package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		usage()
	}
	fname := os.Args[1]
	cmdStr := os.Args[2:]

	inputData, readErr := ioutil.ReadFile(fname)
	if readErr != nil {
		fmt.Println("Unable to open", fname, ":", readErr)
		os.Exit(1)
	}

	var output bytes.Buffer

	var cmd = exec.Command(cmdStr[0], cmdStr[1:]...)
	cmd.Stdin  = bytes.NewReader(inputData)
	cmd.Stdout = &output

	cmdErr := cmd.Run()

	if cmdErr != nil {
		fmt.Println("Error running command:", cmdErr)
		os.Exit(1)
	}


	// TODO: stat the file first to get these
	writeErr := ioutil.WriteFile(fname, output.Bytes(), 0644)
	if writeErr != nil {
		fmt.Println("Error writing to file:", writeErr)
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("usage: inplace <file> <cmd> [args...]")
	os.Exit(1)
}
