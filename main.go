package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dnnnvx/periph-playground/playground"
)

func main() {

	fmt.Printf("Starting...\n")

	if len(os.Args) > 0 {
		if len(os.Args[1]) == 0 {
			log.Fatal("No exercise selected!")
		}

		exercise := os.Args[1]
		fmt.Printf("Exercise selected: %s\n", exercise)

		switch exercise {
		case "blink":
			gpioNum := 17
			if len(os.Args) > 2 {
				gpioNum, err := strconv.Atoi(os.Args[2])
				if err != nil || gpioNum <= 0 || gpioNum > 20 {
					gpioNum = 17
				}
			}
			playground.Blink(gpioNum)
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

}
