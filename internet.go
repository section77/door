package main

import (
	"log"

	"github.com/section77/door/gpio"
)

func enableInternet() {
	log.Println("enable intenet")
	gpio.Write(gpio.Pin15, gpio.Low)
}

func disableInternet() {
	log.Println("disable internet")
	gpio.Write(gpio.Pin15, gpio.High)
}
