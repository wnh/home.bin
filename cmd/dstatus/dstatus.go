// +build linux

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
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
	for t := time.Tick(1 * time.Second); ; <-t {
		status := doStatus()
		xproto.ChangeProperty(x, xproto.PropModeReplace, root, xproto.AtomWmName,
		xproto.AtomString, 8, uint32(len(status)), []byte(status))
	}
}

func doStatus() string {
	t := timeStr()
	b :=  battery()
	v :=  getVolume()

	return fmt.Sprintf("v%v  b%v  %v", v, b, t)
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
	} else if status == "Unknown" {
		symbol = "?"
	}

	return fmt.Sprintf("%v%v", cap, symbol)
}

func timeStr() string {
	d := time.Now()
	dateStr := d.Format("Mon 2006/01/02 15:04")
	return dateStr
}

func readString(fname string) string {
	str, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("error reading ", fname, ":", err)
		os.Exit(1)
	}
	return string(bytes.TrimSpace(str))
}

func getVolume() string {
	var outbuf bytes.Buffer
	cmd := exec.Command("pacmd", "list-sinks")
	cmd.Stdout = &outbuf
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error getting volume:", err)
	}
	out := outbuf.String()
	outlines :=  strings.Split(out, "\n")
	for _, l := range outlines {
		clean := strings.TrimSpace(l)
		if strings.HasPrefix(clean, "volume:") {
			parts := strings.Split(clean, "/")
			cleanVol := strings.TrimSpace(parts[1])
			return cleanVol
		}
	}
	return ""
}
