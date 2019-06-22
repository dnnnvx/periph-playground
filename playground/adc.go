package playground

import (
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host/rpi"
)

func adc() {

	p := rpi.P1_27
	t := time.NewTicker(500 * time.Millisecond)
	for l := gpio.Low; ; l = !l {
		p.Out(l)
		<-t.C
	}

}
