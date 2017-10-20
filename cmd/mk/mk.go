package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const FILE string = "Makefile"
const PROG string = "make"

/* TODO:
add stopping at homedir
configurable filename for aliases?
*/
func main() {
	checkDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for {
		//fmt.Println("Checking:", checkDir)
		if existsAtPath(checkDir) {
			//fmt.Println("FOUND IT in", checkDir)
			os.Exit(runAt(checkDir))
		} else {
			newdir := filepath.Dir(checkDir)
			//fmt.Println("Moving to:", newdir)
			checkDir = newdir
		}
	}
}

func runAt(dir string) int {
	cmd := exec.Command(PROG, os.Args[1:]...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("exit with errr %v\n", err)
		return 1
	} else {
		return 0
	}
}

func existsAtPath(dir string) bool {
	path := filepath.Join(dir, FILE)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	//fmt.Println("found:", path)
	return true
}
