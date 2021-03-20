package playground

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

// Ledbar3pins waterfall ledbar witho only 3 pins
func Ledbar3pins() {

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

	pData.Out(false)
	pLatch.Out(false)
	pClock.Out(false)

	for {
		var i, x uint8

		x = 0x01

		for i = 0; i < 8; i++ {
			pLatch.Out(false)
			shiftOut(pData, pClock, x, false)
			pLatch.Out(true)
			x <<= 1
			time.Sleep(time.Millisecond)
		}

		x = 0x80

		for i = 0; i < 8; i++ {
			pLatch.Out(true)
			shiftOut(pData, pClock, x, false)
			pLatch.Out(false)
			x >>= 1
			time.Sleep(time.Millisecond)
		}
	}
}

func shiftOut(pData gpio.PinOut, pClock gpio.PinOut, x uint8, inverse bool) {
	var i uint8
	for i = 0; i < 8; i++ {
		pClock.Out(false)
		if inverse == false {
			pData.Out((0x01 & (x >> i)) == 0x01)
			time.Sleep(time.Millisecond)
		} else {
			pData.Out((0x80 & (x << i)) == 0x80)
			time.Sleep(time.Millisecond)
		}
		pClock.Out(true)
		time.Sleep(time.Millisecond)
	}
}
