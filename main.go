package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("webapp")))
	http.HandleFunc("/toggle", toggleHndl)
	http.HandleFunc("/open", openHndl)
	http.HandleFunc("/close", closeHndl)
	http.HandleFunc("/state", stateHndl)
	http.HandleFunc("/history", historyHndl)
	fmt.Println("listen on 127.0.0.1:8000")
	fmt.Println(http.ListenAndServe("127.0.0.1:8000", nil))
}

func toggleHndl(w http.ResponseWriter, r *http.Request) {
	lockToggle()
}

func openHndl(w http.ResponseWriter, r *http.Request) {
	lockOpen()
}

func closeHndl(w http.ResponseWriter, r *http.Request) {
	lockClose()
}

func stateHndl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "state: %s", state)
}

func historyHndl(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(history)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
