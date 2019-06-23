package main

import (
	"fmt"
	"log"
	"rpi-test/playground"

	"periph.io/x/periph/host"
)

func main() {

	fmt.Printf("Starting...\n")

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	playground.Segments()
}
