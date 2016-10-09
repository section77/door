// entry point for the application
package main

import (
	"flag"
	"time"
)

var dl DoorLock

func main() {

	// args
	var printHelp = flag.Bool("h", false, "print this help")
	var ip = flag.String("ip", "127.0.0.1", "bind the embedded webserver to the given ip address")
	var port = flag.Int("port", 8000, "bind the embedded webserver to the given port")

	var keepInternet = flag.Bool("keep-internet", false, "keep the internet on")
	var maxHistoryEntries = flag.Int("max-history", 50, "keep at most 'n' entries in the history")

	var btnMinPushDur = flag.Duration("btn-min-push", 500*time.Millisecond, "minimum push duration for the button to react")
	var btnLockDelayDur = flag.Duration("btn-lock-delay", 10*time.Second, "delay duration to lock the door when triggered per button")

	flag.Parse()

	// action
	if *printHelp {
		flag.Usage()
	} else {
		dl = NewDoorLock(*keepInternet, *maxHistoryEntries)

		// start button listener
		startButtonObserverLoop(*btnMinPushDur, *btnLockDelayDur)

		// startup the webserver
		startWebapp(*ip, *port)
	}
}
