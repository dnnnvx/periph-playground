package playground

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/gpio"
)

// Pir exercise for the passive infrared sensor
func Pir(p gpio.PinIO) {

	if err := p.In(gpio.PullNoChange, gpio.RisingEdge); err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting %s (%s)...\n", p, p.Read())

	// Wait for edges as detected by the hardware
	for p.WaitForEdge(-1) {
		if p.Read() == gpio.High {
			fmt.Printf("You moved!\n")
		}
	}
}
