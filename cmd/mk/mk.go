package main

import (
	"fmt"
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
		fmt.Println("Error getting working directory:", err)
		os.Exit(1)
	}
	for {
		if existsAtPath(checkDir) {
			os.Exit(runAt(checkDir))
		} else if checkDir == "/" {
			fmt.Println("Unable to find", FILE)
			os.Exit(1)
		} else {
			newdir := filepath.Dir(checkDir)
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
		fmt.Println(PROG, "exited with error", err)
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
	return true
}
