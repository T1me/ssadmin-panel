package main

import (
	"log"
)

func checkErr(err error, funcName string) {
	if err != nil {
		log.Fatal(funcName, ":", err)
		panic(err)
	}
}
