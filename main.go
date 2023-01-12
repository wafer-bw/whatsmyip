package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/wafer-bw/whatsmyip/api"
)

func getEnv(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}

func main() {
	p := getEnv("HTTP_PORT", "8000")
	s := &http.Server{
		Addr:         net.JoinHostPort("0.0.0.0", p),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  1 * time.Minute,
		Handler:      api.GetRouter(),
	}
	log.Printf("Listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
