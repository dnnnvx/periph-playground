package playground

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

// Segments prints A-to-F every 500ms
func Segments() {

	var pData gpio.PinOut
	var pLatch gpio.PinOut
	var pClock gpio.PinOut

	num := [16]uint{0xc0, 0xf9, 0xa4, 0xb0, 0x99, 0x92, 0x82, 0xf8, 0x80, 0x90, 0x88, 0x83, 0xc6, 0xa1, 0x86, 0x8e}

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
		for i := 0; i < len(num); i++ {
			pLatch.Out(false)
			segmentShift(pData, pClock, num[i], true)
			pLatch.Out(true)
			time.Sleep(time.Millisecond * 500)
		}
		for i := 0; i < len(num); i++ {
			pLatch.Out(true)
			// If you want to display the decimal point,
			// make the highest bit of each array become 0,
			// which can be implemented easily by num[i]&0x7f
			segmentShift(pData, pClock, num[i]&0x7f, true)
			pLatch.Out(false)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func segmentShift(pData gpio.PinOut, pClock gpio.PinOut, val uint, inverse bool) {
	var i uint8
	for i = 0; i < 8; i++ {
		pClock.Out(false)
		if inverse == false {
			pData.Out((0x01 & (val >> i)) == 0x01)
			time.Sleep(time.Microsecond * 10)
		} else {
			pData.Out((0x80 & (val << i)) == 0x80)
			time.Sleep(time.Microsecond * 10)
		}
		pClock.Out(true)
		time.Sleep(time.Microsecond * 10)
	}
}
