package playground

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

// ButtonLed lightup a LED by button pressing
func ButtonLed() {

	// Lookup a pin by its number:
	p := gpioreg.ByName("18")
	if p == nil {
		log.Fatal("Failed to find 18")
	}

	fmt.Printf("%s: %s\n", p, p.Function())

	// Set it as input, with an internal pull down resistor:
	if err := p.In(gpio.PullDown, gpio.BothEdges); err != nil {
		log.Fatal(err)
	}

	led := gpioreg.ByName("17")

	// Wait for edges as detected by the hardware, and print the value read:
	for {
		p.WaitForEdge(-1)
		btnStat := p.Read()
		led.Out(!btnStat)
		// fmt.Printf("Pressed!\n")
	}
}
