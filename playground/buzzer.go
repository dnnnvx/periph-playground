package playground

import (
	"log"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/physic"
)

// Buzzer sound buzz on button pressed
func Buzzer() {

	// Lookup a pin by its number:
	button := gpioreg.ByName("18")
	if button == nil {
		log.Fatal("Failed to find button (pin 18)")
	}
	// Lookup a pin by its number:
	transistor := gpioreg.ByName("17")
	if transistor == nil {
		log.Fatal("Failed to find button (pin 18)")
	}

	// Set it as input, with an internal pull down resistor:
	if err := button.In(gpio.PullDown, gpio.BothEdges); err != nil {
		log.Fatal(err)
	}

	if err := transistor.Halt(); err != nil {
		log.Fatal(err)
	}

	// Wait for edges as detected by the hardware, and print the value read:
	for {
		button.WaitForEdge(-1)
		btnStat := button.Read()
		if btnStat == gpio.Low {
			transistor.PWM(gpio.DutyHalf, 660*physic.Hertz)
		} else {
			transistor.Halt()
		}
	}
}
