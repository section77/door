package main

import (
	"log"
	"time"

	"github.com/section77/door/gpio"
)

func startButtonObserverLoop(btnMinPushDur time.Duration, btnLockDelayDur time.Duration) {

	go func() {
		for {
			if isPushed, duration, err := gpio.BlockWhileIsLow(gpio.Pin11); err != nil {
				log.Printf("unable to query gpio low state: %s\n", err.Error())
			} else {
				if isPushed {
					log.Printf("button was pushed for: '%s' - ", duration)
					switch {
					case duration.Seconds() < btnMinPushDur.Seconds():
						// ignore - to short
						log.Printf("less than the configured minimum from: %s - ignore", btnMinPushDur)
					default:
						switch dl.State {
						case Locked:
							dl.Unlock()
						case Unknown, Unlocked:
							log.Printf("lock door after %s\n", btnLockDelayDur)
							time.Sleep(btnLockDelayDur)
							dl.Lock()
						}
					}
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
}
