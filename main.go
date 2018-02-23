package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	USERNAME      string
	PASSWORD      string
	TMPDIR        string = "/tmp/ssadmin-panel"
	SSUSERSPATH   string = "/tmp/ssadmin-panel/ssusers"
	SSTRAFFICPATH string = "/tmp/ssadmin-panel/sstraffic"
)

func update(second int) {
	exec.Command("mkdir " + TMPDIR)
	showCommand := "ssadmin show > " + SSUSERSPATH
	showpwCommand := "ssadmin showpw > " + SSTRAFFICPATH
	n := time.Duration(second)
	for {
		exec.Command(showCommand)
		exec.Command(showpwCommand)
		time.Sleep(n * time.Second)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ssadmin-panel username password")
		return
	}
	USERNAME = os.Args[1]
	PASSWORD = os.Args[2]
	go update(300)
	DefineListenAndServe()
}
