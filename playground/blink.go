package playground

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
)

// Blink LED every second
func Blink(p gpio.PinIO) {
	fmt.Printf("Starting...\n")
	t := time.NewTicker(1000 * time.Millisecond)
	for l := gpio.Low; ; l = !l {
		if err := p.Out(l); err != nil {
			log.Fatal(err)
		}
		<-t.C
	}
}
