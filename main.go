package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	state   = "keine ahnung!"
	history [][]string
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("webapp")))
	http.HandleFunc("/open", openHndl)
	http.HandleFunc("/close", closeHndl)
	http.HandleFunc("/state", stateHndl)
	http.HandleFunc("/history", historyHndl)
	fmt.Println("listen on :80")
	fmt.Println(http.ListenAndServe(":80", nil))
}

func openHndl(w http.ResponseWriter, r *http.Request) {
	state = "open"
	addToHistory("open event")
	dur, _ := time.ParseDuration("500us")
	servo(dur)
}

func closeHndl(w http.ResponseWriter, r *http.Request) {
	state = "close"
	addToHistory("close event")
	dur, _ := time.ParseDuration("2ms")
	servo(dur)
}

func stateHndl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "state: %s", state)
}

func historyHndl(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(history)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func servo(duration time.Duration) {
	GPIOOpen()
	time.Sleep(500 * time.Millisecond)

	GPIOLow()
	for i := 0; i < 50; i++ {
		GPIOHigh()
		time.Sleep(duration)

		GPIOLow()
		time.Sleep(10 * time.Millisecond)
	}

	GPIOClose()
}

func addToHistory(msg string) {
	ts := time.Now().Format("02.01 15:04")
	history = append(history, []string{ts, msg})
}
