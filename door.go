package main

import (
	"fmt"
	"log"
	"time"

	"github.com/section77/door/gpio"
)

const (
	// Unknown represents the state when the app starts up
	Unknown State = iota
	// Locked represents the state when the door is locked
	Locked
	// Unlocked represents the state when the door is unlocked
	Unlocked
)

// DoorLock represents the door lock and provides functions to lock / unlock
type DoorLock struct {
	State             State
	KeepInternet      bool
	MaxHistoryEntries int
	History           [][]string
}

func NewDoorLock(keepInternet bool, maxHistoryEntries int) DoorLock {
	return DoorLock{
		State:             Unknown,
		KeepInternet:      keepInternet,
		MaxHistoryEntries: maxHistoryEntries,
	}
}

func (dl DoorLock) String() string {
	return fmt.Sprintf("DoorLock(State: %s)", dl.State)
}

func (dl *DoorLock) Lock() {
	if dl.State != Locked {
		dl.addToHistory("lock")
		if err := gpio.Pwm(gpio.Pin12, 80); err != nil {
			log.Printf("unable to lock door: %s\n", err.Error())
		} else {
			dl.State = Locked
			if !dl.KeepInternet {
				disableInternet()
			}
		}
	} else {
		log.Println("already locked")
	}
}

func (dl *DoorLock) Unlock() {
	if dl.State != Unlocked {
		dl.addToHistory("unlock")
		if err := gpio.Pwm(gpio.Pin12, 20); err != nil {
			log.Printf("unable to lock door: %s\n", err.Error())
		} else {
			dl.State = Unlocked
			if !dl.KeepInternet {
				enableInternet()
			}
		}
	} else {
		log.Println("already unlocked")
	}
}

func (dl *DoorLock) toggleLock() {
	log.Println("toggle")
	switch dl.State {
	case Locked:
		dl.Unlock()
	case Unlocked:
		dl.Lock()
	default:
		log.Println("toggle event, but unknow door state - lock the door")
		dl.Lock()
	}
}

func (dl *DoorLock) addToHistory(state string) {
	log.Println(state)
	ts := time.Now().Format("02.01 15:04:05")
	dl.History = append(dl.History, []string{ts, state})

	// delete old entries when history grows
	count := len(dl.History)
	if count >= dl.MaxHistoryEntries {
		dl.History = dl.History[count-dl.MaxHistoryEntries:]
	}
}

type State int

func (s State) String() string {
	str := "- unknow -"
	switch s {
	case Locked:
		str = "locked"
	case Unlocked:
		str = "unlocked"
	}
	return str
}
