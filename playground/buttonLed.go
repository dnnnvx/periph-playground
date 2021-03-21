package playground

import (
	"log"

	"periph.io/x/conn/v3/gpio"
)

// ButtonLed lightup a LED by button pressing
func ButtonLed(pBtn gpio.PinIO, pLed gpio.PinIO) {

	if err := pBtn.In(gpio.PullDown, gpio.BothEdges); err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting %s (%s)...\n", pBtn, pBtn.Read())

	// Wait for edges as detected by the hardware, and print the value read:
	for pBtn.WaitForEdge(-1) {
		btnStat := pBtn.Read()
		pLed.Out(btnStat)
		log.Printf("%s went %s\n", pBtn, gpio.High)
	}
}
