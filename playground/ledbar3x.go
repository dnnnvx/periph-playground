package playground

import (
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

func Ledbar3x() {

	// Initialize pins
	var pData gpio.PinOut
	var pLatch gpio.PinOut
	var pClock gpio.PinOut

	pData = gpioreg.ByName("17")
	pLatch = gpioreg.ByName("27")
	pClock = gpioreg.ByName("22")
	if pData == nil || pClock == nil || pLatch == nil {
		log.Fatal("Failed to load pins")
	}
	// Set them to low
	pData.Out(gpio.Low)
	pLatch.Out(gpio.Low)
	pClock.Out(gpio.Low)

	for {
		var i, x uint8
		var inverse bool

		x = 0x01
		i = 0x00

		for ; i < 0x80; i++ {
			pLatch.Out(gpio.Low)
			shiftOut(pData, pClock, x, inverse)
			pLatch.Out(gpio.High)
			x <<= 1
			time.Sleep(time.Second)
		}

		x = 0x80
		i = 0x00
		inverse = true

		for ; i < 0x80; i++ {
			pLatch.Out(gpio.Low)
			shiftOut(pData, pClock, x, inverse)
			pLatch.Out(gpio.High)
			x >>= 1
			time.Sleep(time.Second)
		}
	}
}

func shiftOut(pData gpio.PinOut, pClock gpio.PinOut, x uint8, inverse bool) {
	var i uint8
	for i = 0; i < 8; i++ {
		pClock.Out(gpio.Low)
		if !inverse {
			if w := x >> i; w == 0x01 {
				pData.Out(gpio.Low)
			} else {
				pData.Out(gpio.High)
			}
			time.Sleep(time.Second)
		} else {
			if w := x >> i; w == 0x80 {
				pData.Out(gpio.Low)
			} else {
				pData.Out(gpio.High)
			}
			time.Sleep(time.Second)
		}
		pClock.Out(gpio.High)
		time.Sleep(time.Second)
	}

}
