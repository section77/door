package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
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
	addToHistory("open")
	pwm(20)
}

func closeHndl(w http.ResponseWriter, r *http.Request) {
	addToHistory("close")
	pwm(80)
}

func stateHndl(w http.ResponseWriter, r *http.Request) {
	state := "keine ahnung!"
	if len(history) > 0 {
		// pfusch!!
		state = history[len(history)-1][1]
	}
	fmt.Fprintf(w, "state: %s", state)
}

func historyHndl(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(history)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func addToHistory(msg string) {
	ts := time.Now().Format("02.01 15:04")
	history = append(history, []string{ts, msg})
}
