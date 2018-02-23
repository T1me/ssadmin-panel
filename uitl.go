package main

import (
	"log"
)

func ThrownErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CatchErr() {
	if err := recover(); err != nil {
		log.Fatal(err)
	}
}
