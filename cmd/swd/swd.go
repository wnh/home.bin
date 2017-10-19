package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("[ERR] %v\n", err)
		return
	}

	home := os.Getenv("HOME")
	if strings.HasPrefix(dir, home) {
		dir = "~" + strings.TrimPrefix(dir, home)
	}

	dirs := strings.Split(dir, "/")
	out := ""
	for i, d := range(dirs) {
	  if d == "" {
		continue
	  } else if d == "~" {
		out += d + "/"
	  } else if i == (len(dirs) - 1) {
		out += d
		break
	  } else {
		// TODO(wnh); UUUUUUGLY
		out += string([]rune(d)[0])
		out += "/"
	  }
	}
	fmt.Printf("%s", out)
}
