package playground

import (
	"log"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/physic"
)

// Buzzer sound buzz on button pressed
func Buzzer(pBtn gpio.PinIO, pTran gpio.PinIO) {

	// Set it as input, with an internal pull down resistor
	if err := pBtn.In(gpio.PullDown, gpio.BothEdges); err != nil {
		log.Fatal(err)
	}

	var sound physic.Frequency = 200

	// Wait for edges as detected by the hardware, and print the value read:
	for pBtn.WaitForEdge(-1) {
		btnStat := pBtn.Read()
		if btnStat == gpio.Low {
			if err := pTran.Halt(); err != nil {
				log.Fatal(err)
			}
		} else {
			if pTran.Read() == gpio.Low {
				if err := pTran.PWM(gpio.DutyHalf, sound*physic.Hertz); err != nil {
					log.Fatal(err)
				}
			}
			if sound == 800 {
				sound = 200
			} else {
				sound += 200
			}
		}
	}
}
