package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/wafer-bw/whatsmyip/spec"
	"google.golang.org/protobuf/proto"
)

func resolver(request *http.Request) *spec.IPReply {
	ip := request.Header.Get("x-forwarded-for")
	if ip == "" {
		ip = strings.Split(request.RemoteAddr, ":")[0]
	}
	return &spec.IPReply{Ip: ip}
}

func marshaller(request *http.Request, reply *spec.IPReply) (body []byte, err error) {
	switch request.Header.Get("accept") {
	case "application/protobuf":
		return proto.Marshal(reply)
	case "application/json":
		return json.Marshal(reply)
	default:
		return []byte(reply.Ip), nil
	}
}

// Handler responds with the IP address of the request
func Handler(writer http.ResponseWriter, request *http.Request) {
	log.Println(*request)
	body, err := marshaller(request, resolver(request))
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), 500)
	}
	writer.Write(body)
}

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler).Methods(http.MethodGet)
	return r
}

func getEnv(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}

func main() {
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
