package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/wafer-bw/whatsmyip/spec"
	"google.golang.org/protobuf/proto"
)

// GetRouter returns the router for the API
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler).Methods(http.MethodGet)
	return r
}

func resolver(request *http.Request) *spec.IPReply {
	ip := request.Header.Get("x-forwarded-for")
	if ip == "" {
		ip = strings.Split(request.RemoteAddr, ":")[0]
	}
	return &spec.IPReply{Ip: ip}
}

func marshaller(writer http.ResponseWriter, request *http.Request, reply *spec.IPReply) (body []byte, err error) {
	switch request.Header.Get("accept") {
	case "application/protobuf":
		writer.Header().Set("Content-Type", "application/protobuf")
		return proto.Marshal(reply)
	case "application/json":
		writer.Header().Set("Content-Type", "application/json")
		return json.Marshal(reply)
	default:
		writer.Header().Set("Content-Type", "text/plain")
		return []byte(reply.Ip), nil
	}
}

// Handler responds with the IP address of the request
func Handler(writer http.ResponseWriter, request *http.Request) {
	log.Println(*request)
	body, err := marshaller(writer, request, resolver(request))
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), 500)
	}
	writer.Write(body)
}
