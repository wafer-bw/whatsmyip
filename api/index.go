package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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
	body, err := marshaller(request, resolver(request))
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), 500)
	}
	writer.Write(body)
}
