package playground

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
)

// Segments 8 segment led output
func Segments(pData gpio.PinIO, pLatch gpio.PinIO, pClock gpio.PinIO) {

	num := [16]uint{
		0xc0, // binary: 11000000, decimal: 192 (displayed 0)
		0xf9, // binary: 11111001, decimal: 249 (displayed 1)
		0xa4, // binary: 10100100, decimal: 164 (displayed 2)
		0xb0, // binary: 10110000, decimal: 176 (displayed 3)
		0x99, // binary: 10011001, decimal: 153 (displayed 4)
		0x92, // binary: 10010010, decimal: 146 (displayed 5)
		0x82, // binary: 10000010, decimal: 130 (displayed 6)
		0xf8, // binary: 11111000, decimal: 248 (displayed 7)
		0x80, // binary: 10000000, decimal: 128 (displayed 8)
		0x90, // binary: 10010000, decimal: 144 (displayed 9)
		0x88, // binary: 10001000, decimal: 136 (displayed a)
		0x83, // binary: 10000011, decimal: 131 (displayed b)
		0xc6, // binary: 11000110, decimal: 198 (displayed c)
		0xa1, // binary: 10100001, decimal: 161 (displayed d)
		0x86, // binary: 10000110, decimal: 134 (displayed e)
		0x8e, // binary: 10001110, decimal: 142 (displayed f)
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
			pLatch.Out(false)
			// now with the decimal point:
			// make the highest bit of each array become 0,
			// which can be implemented easily by num[i]&0x7f
			segmentShift(pData, pClock, num[i]&0x7f, true)
			pLatch.Out(true)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func segmentShift(pData gpio.PinOut, pClock gpio.PinOut, val uint, inverse bool) {
	var i uint8
	// i = 0 (binary: 00000000) -> 00000001 & (val >> 00000000) == 00000001
	// i = 1 (binary: 00000001) -> 00000001 & (val >> 00000001) == 00000001
	// i = 2 (binary: 00000010) -> 00000001 & (val >> 00000010) == 00000001
	// i = 3 (binary: 00000011) -> 00000001 & (val >> 00000011) == 00000001
	// i = 4 (binary: 00000100) -> 00000001 & (val >> 00000100) == 00000001
	// i = 5 (binary: 00000101) -> 00000001 & (val >> 00000101) == 00000001
	// i = 6 (binary: 00000110) -> 00000001 & (val >> 00000110) == 00000001
	// i = 7 (binary: 00000111) -> 00000001 & (val >> 00000111) == 00000001
	for i = 0; i < 8; i++ {

		switch i {
		case 0: // i = 0 (binary: 00000000) -> 00000001 & (val >> 00000000) == 00000001
			log.Printf(" i = 0 (binary: 00000000) -> 00000001 & (%b << 00000000) == 00000001 | %t", val, (0x01&(val<<i) == 0x01))
			break
		case 1: // i = 1 (binary: 00000001) -> 00000001 & (val >> 00000001) == 00000001
			log.Printf(" i = 1 (binary: 00000001) -> 00000001 & (%b << 00000001) == 00000001 | %t", val, (0x01&(val<<i) == 0x01))
			break
		case 2: // i = 2 (binary: 00000010) -> 00000001 & (val >> 00000010) == 00000001
			log.Printf(" i = 2 (binary: 00000010) -> 00000001 & (%b << 00000010) == 00000001 | %t", val, (0x01&(val<<i) == 0x01))
			break
		case 3: // i = 3 (binary: 00000011) -> 00000001 & (val >> 00000011) == 00000001
			log.Printf(" i = 3 (binary: 00000011) -> 00000001 & (%b << 00000011) == 00000001 | %t", val, (0x01&(val<<i) == 0x01))
			break
		case 4: // i = 4 (binary: 00000100) -> 00000001 & (val >> 00000100) == 00000001
			log.Printf(" i = 4 (binary: 00000100) -> 00000001 & (%b << 00000100) == 00000001 | %t", val, (0x01&(val<<i) == 0x01))
			break
		case 5: // i = 5 (binary: 00000101) -> 00000001 & (val >> 00000101) == 00000001
			log.Printf(" i = 5 (binary: 00000101) -> 00000001 & (%b << 00000101) == 00000001 | %t", val, (0x01&(val<<i) == 0x01))
			break
		case 6: // i = 6 (binary: 00000110) -> 00000001 & (val >> 00000110) == 00000001
			log.Printf(" i = 6 (binary: 00000110) -> 00000001 & (%b << 00000110) == 00000001 | %t", val, (0x01&(val<<i) == 0x01))
			break
		case 7: // i = 7 (binary: 00000111) -> 00000001 & (val >> 00000111) == 00000001
			log.Printf(" i = 7 (binary: 00000111) -> 00000001 & (%b << 00000111) == 00000001 | %t", val, (0x01&(val<<i) == 0x01))
			break
		}

		pClock.Out(false)
		if inverse == false {
			// hex: 0x01, binary: 1 (00000001), decimal: 1
			pData.Out((0x01 & (val >> i)) == 0x01)
			// time.Sleep(time.Microsecond * 10)
		} else {
			// hex: 0x80, binary: 10000000, decimal: 128
			pData.Out((0x80 & (val << i)) == 0x80)
			// time.Sleep(time.Microsecond * 10)
		}
		pClock.Out(true)
		// time.Sleep(time.Microsecond * 10)
	}
}
