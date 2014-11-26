package main

import "github.com/stianeikeland/go-rpio"

var (
	pin = rpio.Pin(4)
)

func GPIOOpen() error {
	if err := rpio.Open(); err != nil {
		return err
	}
	pin.Output()
	return nil
}

func GPIOClose() {
	rpio.Close()
}

func GPIOHigh() {
	pin.High()
}

func GPIOLow() {
	pin.Low()
}

func GPIOFlip() {
	pin.Toggle()
}
