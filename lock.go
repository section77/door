package main

import (
	"log"
	"time"
)

var (
	state   State
	history [][]string
)

type State int

const (
	Unknow State = iota
	Open
	Close
)

func (s State) String() string {
	str := "- unknown -"
	switch s {
	case Open:
		str = "open"
	case Close:
		str = "close"
	}
	return str
}

func lockOpen() {
	state = Open
	addToHistory("open event")
	enableUMTSStick(true)
	pwm(20)
}

func lockClose() {
	state = Close
	addToHistory("close event")
	enableUMTSStick(false)
	pwm(80)
}

func lockToggle() {
	switch state {
	case Open:
		lockClose()
	case Close:
		lockOpen()
	default:
		log.Printf("WARNUNG: aktueller status: '%s' - mache zu!\n", state)
		lockClose()
	}
}

func addToHistory(msg string) {
	ts := time.Now().Format("02.01 15:04")
	history = append(history, []string{ts, msg})
}
