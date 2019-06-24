package playground

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host/rpi"
)

// Ledbar waterfall lights with the ledbar (10 pins)
func Ledbar() {

	// Lookup pins
	p1 := gpioreg.ByName("18")
	if p1 == nil {
		log.Fatal("Failed 1")
	}
	p2 := gpioreg.ByName("17")
	if p2 == nil {
		log.Fatal("Failed 2")
	}
	p3 := gpioreg.ByName("22")
	if p3 == nil {
		log.Fatal("Failed 3")
	}
	p4 := gpioreg.ByName("27")
	if p4 == nil {
		log.Fatal("Failed 4")
	}
	p5 := gpioreg.ByName("23")
	if p5 == nil {
		log.Fatal("Failed  5")
	}
	p6 := gpioreg.ByName("24")
	if p6 == nil {
		log.Fatal("Failed 6")
	}
	p7 := gpioreg.ByName("25")
	if p7 == nil {
		log.Fatal("Failed 7")
	}
	p8 := rpi.P1_3
	if p8 == nil {
		log.Fatal("Failed 8")
	}
	p9 := rpi.P1_5
	if p9 == nil {
		log.Fatal("Failed 9")
	}
	p0 := rpi.P1_24
	if p0 == nil {
		log.Fatal("Failed 0")
	}

	fmt.Printf("%s: %s\n", p1, p1.Function())

	// Set it as input, with an internal pull down resistor:
	if err := p0.Out(false); err != nil {
		log.Fatal(err)
	}

	t := time.NewTicker(500 * time.Millisecond)

	i := 0
	pre := 1

	for {
		p1.Out(gpio.High)
		p2.Out(gpio.High)
		p3.Out(gpio.High)
		p4.Out(gpio.High)
		p5.Out(gpio.High)
		p6.Out(gpio.High)
		p7.Out(gpio.High)
		p8.Out(gpio.High)
		p9.Out(gpio.High)
		p0.Out(gpio.High)
		switch i {
		case 0:
			p1.Out(gpio.Low)
			break
		case 1:
			p2.Out(gpio.Low)
			break
		case 2:
			p3.Out(gpio.Low)
			break
		case 3:
			p4.Out(gpio.Low)
			break
		case 4:
			p5.Out(gpio.Low)
			break
		case 5:
			p6.Out(gpio.Low)
			break
		case 6:
			p7.Out(gpio.Low)
			break
		case 7:
			p8.Out(gpio.Low)
			break
		case 8:
			p9.Out(gpio.Low)
			break
		case 9:
			p0.Out(gpio.Low)
			break
		}
		if pre == 1 {
			i++
		} else if pre == 0 {
			i--
		}
		if i == 9 {
			pre = 0
		} else if i == 0 {
			pre = 1
		}
		<-t.C
	}
}
