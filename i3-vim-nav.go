package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/proxypoke/i3ipc"
	"github.com/vbrown608/xdo-go"
)

func main() {
	dir := string(os.Args[1])

	if match, _ := regexp.MatchString(`h|j|k|l`, dir); !match {
		fmt.Println("must have an argument h j k or l")
		return
	}

	if windowIsVim() {
		keycmd := exec.Command("xdotool", "key", "--clearmodifiers", "Escape+control+"+dir)
		out, _ := keycmd.Output()
		if len(out) > 0 {
			fmt.Println(out)
		}
	} else {
		conn, err := i3ipc.GetIPCSocket()
		if err != nil {
			fmt.Println("could not connect to i3")
			return
		}
		m := make(map[string]string)
		m["h"] = "left"
		m["j"] = "down"
		m["k"] = "up"
		m["l"] = "right"
		conn.Command("focus " + m[dir])
	}
}

func windowIsVim() bool {
	xdot := xdo.NewXdo()
	window, err := xdot.GetActiveWindow()
	if err != nil {
		return false
	}

	name := strings.ToLower(window.GetName())
	r, _ := regexp.Compile(`\bn?v(im)?$`)
	return r.MatchString(name)
}
