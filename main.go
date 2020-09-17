package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	r.HandleFunc("/", api.Handler).Methods(http.MethodGet)
	p := getEnv("HTTP_PORT", "80")
	s := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", p),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  1 * time.Minute,
		Handler:      getRouter(),
	}
	log.Printf("Listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
