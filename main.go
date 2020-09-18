package main

import (
	"log"
	"net/http"
	"time"

	"github.com/wafer-bw/whatsmyip/api"
)

func main() {
	s := &http.Server{
		Addr:         "0.0.0.0:8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  1 * time.Minute,
		Handler:      api.GetRouter(),
	}
	log.Printf("Listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
