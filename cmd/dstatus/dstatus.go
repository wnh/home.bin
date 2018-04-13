// +build linux

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

func main() {
	x, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	root := xproto.Setup(x).DefaultScreen(x).Root
	for t := time.Tick(5 * time.Second); ; <-t {
		status := fmt.Sprintf("%v  %v", battery(), timeStr())
		xproto.ChangeProperty(x, xproto.PropModeReplace, root, xproto.AtomWmName,
		xproto.AtomString, 8, uint32(len(status)), []byte(status))
	}
}

func battery() string {
	cap    := readString("/sys/class/power_supply/BAT0/capacity")
	status := readString("/sys/class/power_supply/BAT0/status")
	symbol := ""
	if status == "Discharging" {
		symbol = "-"
	} else if status == "Charging" {
		symbol = "+"
	} else if status == "Full" {
		symbol = "!"
	}

	return fmt.Sprintf("%v%v", symbol, cap)
}

func timeStr() string {
	return time.Now().Format("2006/01/02 15:04")
}

func readString(fname string) string {
	str, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("error reading ", fname, ":", err)
		os.Exit(1)
	}
	return string(bytes.TrimSpace(str))
}
