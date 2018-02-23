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
	SSUSERSPATH   string = "/tmp/ssadmin-panel/ssusers"
	SSTRAFFICPATH string = "/tmp/ssadmin-panel/sstraffic"
)

func update(second int) {
	showCommand := "ssadmin show > " + SSUSERSPATH
	showpwCommand := "ssadmin showpw > " + SSTRAFFICPATH
	for {
		time.Sleep(second * time.Second)
		exec.Command(showCommand)
		exec.Command(showpwCommand)
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
