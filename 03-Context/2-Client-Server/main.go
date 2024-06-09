package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request ended")
	select {
	case <-time.After(5 * time.Second):
		// Print on command line stdout
		log.Println("Request processed successfully")
		// Print on browser
		w.Write([]byte("Request processed successfully"))
	case <-ctx.Done():
		// Print on command line stdout
		log.Println("Request cancelled by client")
	}
}
