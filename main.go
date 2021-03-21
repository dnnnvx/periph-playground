package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dnnnvx/periph-playground/playground"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func main() {

	exercise := start(os.Args)
	argslen := len(os.Args)

	switch exercise {

	case "blink":
		if argslen <= 2 {
			log.Fatal("No GPIO selected!")
		}
		initGPIO(os.Args[2])
		gpioNum := initGPIO(os.Args[2])
		playground.Blink(gpioNum)
		break

	case "buttonLed":
		if argslen <= 3 {
			log.Fatal("No GPIO selected!")
		}
		gpioBtn := initGPIO(os.Args[2])
		gpioLed := initGPIO(os.Args[3])
		playground.ButtonLed(gpioBtn, gpioLed)
		break

	case "pir":
		if argslen <= 2 {
			log.Fatal("No GPIO selected!")
		}
		p := initGPIO(os.Args[2])
		playground.Pir(p)
		break

	case "buzzer":
		if argslen <= 3 {
			log.Fatal("No GPIO selected!")
		}
		gpioBtn := initGPIO(os.Args[2])
		gpioTran := initGPIO(os.Args[3])
		playground.Buzzer(gpioBtn, gpioTran)
		break

	case "segments":
		if argslen <= 4 {
			log.Fatal("No GPIO selected!")
		}
		pData := initGPIO(os.Args[2])
		pLatch := initGPIO(os.Args[3])
		pClock := initGPIO(os.Args[4])
		playground.Segments(pData, pLatch, pClock)
		break

	case "ledbar":
		playground.Ledbar()
		break
	case "ledbar3pins":
		playground.Ledbar3pins()
		break
	case "segments4":
		playground.Segments4digit()
		break
	default:
		log.Fatal("Not implemented yet :(")
	}

}

func start(args []string) string {
	if len(args) <= 1 {
		log.Fatal("No exercise selected!")
	}
	exercise := os.Args[1]
	fmt.Printf("Exercise selected: %s\n", exercise)
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	return exercise
}

func initGPIO(arg string) gpio.PinIO {
	gpioNum, err := strconv.Atoi(arg)
	if err != nil || gpioNum <= 0 || gpioNum > 27 {
		log.Fatal("Wrong GPIO selected!")
	}
	pin := gpioreg.ByName(fmt.Sprint(gpioNum))
	if pin == nil {
		log.Fatalf("Failed to find GPIO %d", gpioNum)
	}
	return pin
}
