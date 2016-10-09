package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// startWebapp starts the embedded webserver and
// register the handlers for the endpoints
func startWebapp(ip string, port int) {
	http.Handle("/", http.FileServer(http.Dir("webapp")))
	http.HandleFunc("/toggle", toggleHndl)
	http.HandleFunc("/open", openHndl)
	http.HandleFunc("/close", closeHndl)
	http.HandleFunc("/state", stateHndl)
	http.HandleFunc("/history", historyHndl)

	addr := fmt.Sprintf("%s:%d", ip, port)
	log.Printf("listen on %s\n", addr)
	log.Println(http.ListenAndServe(addr, nil))
}

func toggleHndl(w http.ResponseWriter, r *http.Request) {
	dl.toggleLock()
}

func openHndl(w http.ResponseWriter, r *http.Request) {
	dl.Unlock()
}

func closeHndl(w http.ResponseWriter, r *http.Request) {
	dl.Lock()
}

func stateHndl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "dl lock: %s", dl.State)
}

func historyHndl(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(dl.History)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
