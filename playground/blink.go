package playground

import (
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func blink() {
	host.Init()
	p := gpioreg.ByName("17")
	t := time.NewTicker(500 * time.Millisecond)
	for l := gpio.Low; ; l = !l {
		p.Out(l)
		<-t.C
	}
}
