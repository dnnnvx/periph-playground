package playground

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

// Blink LED (GPIO17) every seconds
func Blink(gpioNum int) {
	fmt.Printf("Starting...\n")
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	log.Printf("GPIO selected: %d", gpioNum)
	p := gpioreg.ByName(fmt.Sprint(gpioNum))
	t := time.NewTicker(1000 * time.Millisecond)
	for l := gpio.Low; ; l = !l {
		if err := p.Out(l); err != nil {
			log.Fatal(err)
		}
		<-t.C
	}
}
