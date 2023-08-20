package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/wafer-bw/whatsmyip/api"
)

var (
	host         = "0.0.0.0"
	defaultPort  = "8000"
	readTimeout  = 5 * time.Second
	writeTimeout = 1 * time.Second
	idleTimeout  = 1 * time.Minute
	portEnv      = "HTTP_PORT"
)

func main() {
	port := os.Getenv(portEnv)
	if port == "" {
		port = defaultPort
	}

	s := &http.Server{
		Addr:         net.JoinHostPort(host, port),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      api.GetRouter(),
	}

	log.Printf("listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
