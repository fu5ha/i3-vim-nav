package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	xdo "github.com/aep/xdo-go"
	"github.com/proxypoke/i3ipc"
)

func main() {
	dir := string(os.Args[1])

	if match, _ := regexp.MatchString(`h|j|k|l`, dir); !match {
		fmt.Println("must have an argument h j k or l")
		return
	}

	xdot := xdo.NewXdo()
	window := xdot.GetActiveWindow()
	name := strings.ToLower(window.GetName())

	r, _ := regexp.Compile(`\bn?v(im)?$`)

	if r.MatchString(name) {
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
