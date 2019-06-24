package main

import (
	"fmt"
	"log"
	"os"
	"periph-playground/playground"

	"periph.io/x/periph/host"
)

func main() {

	fmt.Printf("Starting...\n")

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	playground.Segments()

	exercise := os.Args[1:]
	if !exercise || len(exercise) == 0 {
		log.Fatal("No exercise selected!")
	}

	switch exercise {
	case "blink":
		playground.Blink()
		break
	case "buttonLed":
		playground.ButtonLed()
		break
	case "buzzer":
		playground.Buzzer()
		break
	case "ledbar":
		playground.Ledbar()
		break
	case "ledbar3pins":
		playground.Ledbar3pins()
		break
	case "segments":
		playground.Segments()
		break
	}
}
