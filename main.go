package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/wafer-bw/whatsmyip/api"
)

const (
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

	router := mux.NewRouter()
	router.HandleFunc("/", api.Handler).Methods(http.MethodGet)

	server := &http.Server{
		Addr:         net.JoinHostPort(host, port),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      router,
	}

	log.Printf("listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
