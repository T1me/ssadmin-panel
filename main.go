package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	USERNAME  string
	PASSWORD  string
	SSUSERS   []byte
	SSTRAFFIC []byte
)

// buffer and update ssusers and sstrafic
func update(second int) {
	cmd := "ssadmin"
	showpwParam := "showpw"
	showParam := "show"
	s := time.Duration(second)
	for {
		term := exec.Command(cmd, showpwParam)
		SSUSERS, _ = term.CombinedOutput()
		term = exec.Command(cmd, showParam)
		SSTRAFFIC, _ = term.CombinedOutput()
		time.Sleep(s * time.Second)
	}
}

func main() {
	defer CatchErr()
	if len(os.Args) != 3 {
		fmt.Println("Usage: ssadmin-panel username password")
		return
	}
	USERNAME = os.Args[1]
	PASSWORD = os.Args[2]
	go update(300)
	DefineListenAndServe()
}
