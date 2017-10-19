package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	out, err := exec.Command("git", "branch", "--list").Output()
	if err != nil {
		return
	}
	lines := strings.Split(string(out), "\n")
	for _, branch := range lines {
		if strings.HasPrefix(branch, "*") {
			fmt.Printf("(%s)", strings.Trim(branch, "* "))
		}
	}
}
