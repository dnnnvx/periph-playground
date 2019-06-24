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

	exercise := os.Args[1]
	if len(exercise) == 0 {
		log.Fatal("No exercise selected!")
	} else {
		fmt.Printf("Exercise selected:%s", exercise)
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
	case "segments4":
		playground.Segments4digit()
		break
	default:
		log.Fatal("Not implemented yet :(")
	}
}
