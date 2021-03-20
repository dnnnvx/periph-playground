package playground

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

var pData gpio.PinOut
var pLatch gpio.PinOut
var pClock gpio.PinOut
var pDigit = [4]gpio.PinOut{}
var num = [10]uint{0xc0, 0xf9, 0xa4, 0xb0, 0x99, 0x92, 0x82, 0xf8, 0x80, 0x90} // 0x88, 0x83, 0xc6, 0xa1, 0x86, 0x8e
var counter = 0

// Segments4digit ...
func Segments4digit() {

	pData = gpioreg.ByName("18")
	pLatch = gpioreg.ByName("23")
	pClock = gpioreg.ByName("24")
	if pData == nil || pClock == nil || pLatch == nil {
		log.Fatal("Failed to load pins")
	}

	pDigit[0] = gpioreg.ByName("17")
	pDigit[1] = gpioreg.ByName("27")
	pDigit[2] = gpioreg.ByName("22")
	pDigit[3] = gpioreg.ByName("10")

	if pDigit[0] == nil || pDigit[1] == nil || pDigit[2] == nil || pDigit[3] == nil {
		log.Fatal("Failed to load pins")
	}

	pData.Out(false)
	pLatch.Out(false)
	pClock.Out(false)

	// Set the pin connected to 7-segment display
	for i := 0; i < 4; i++ {
		pDigit[i].Out(true)
	}

	for {
		display(counter)
		counter++
		time.Sleep(time.Second)
	}
}

func selectDigit(digit int) {
	pDigit[0].Out((digit & 0x08) != 0x08)
	pDigit[1].Out((digit & 0x04) != 0x04)
	pDigit[2].Out((digit & 0x02) != 0x02)
	pDigit[3].Out((digit & 0x01) != 0x01)
}

func segment4dShift(pData gpio.PinOut, pClock gpio.PinOut, val uint, inverse bool) {
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

// Function used to output data for 74HC595
func outData(data uint) {
	pLatch.Out(false)
	segment4dShift(pData, pClock, data, false)
	pLatch.Out(true)
}

// Display function for 7-segment display
func display(dec int) {
	var delays time.Duration = 1
	outData(0xff)
	selectDigit(0x01) // select the first, and display the single digit
	outData(num[dec%10])
	time.Sleep(time.Second * delays)

	outData(0xff)
	selectDigit(0x02) // select the second, and display the tens digit
	outData(num[dec%100/10])
	time.Sleep(time.Second * delays)

	outData(0xff)
	selectDigit(0x04) // select the third, and display the hundreds digit
	outData(num[dec%1000/100])
	time.Sleep(time.Second * delays)

	outData(0xff)
	selectDigit(0x08) // select the fourth, and display the thousands digit
	outData(num[dec%10000/1000])
	time.Sleep(time.Second * delays)
}
