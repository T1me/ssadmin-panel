package main

import (
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
	command := "ssadmin"
	showParams := []string{"show", ">", SSUSERSPATH}
	showpwParams := []string{"showpw", ">", SSTRAFFICPATH}
	ticker := time.NewTicker(time.Second * second)
	for _ = range ticker.C {
		exec.Command(command, showParams)
		exec.Command(command, showpwParams)
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
