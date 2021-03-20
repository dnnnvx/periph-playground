package playground

import (
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3/rpi"
)

func adc() {

	p := rpi.P1_27
	t := time.NewTicker(500 * time.Millisecond)
	for l := gpio.Low; ; l = !l {
		p.Out(l)
		<-t.C
	}

}
