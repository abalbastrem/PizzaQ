package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("HELLO WORLD")
	// send()
	receive()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
