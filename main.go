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
	SSUSERSPATH   string = "ssusers"
	SSTRAFFICPATH string = "sstraffic"
)

func update(second int) {
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
	if len(os.Args) != 4 {
		fmt.Println("Usage: ssadmin-panel username password path")
		return
	}
	USERNAME = os.Args[1]
	PASSWORD = os.Args[2]
	SSUSERSPATH = os.Args[3] + SSUSERSPATH
	SSTRAFFICPATH = os.Args[3] + SSTRAFFICPATH
	go update(300)
	DefineListenAndServe()
}
