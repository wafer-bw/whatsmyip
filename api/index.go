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

// GetRouter returns the router for the API.
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

func respond(w http.ResponseWriter, _ *http.Request, body []byte, err error) {
	switch err {
	case nil:
		_, _ = w.Write(body)
	default:
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func getAcceptHeader(r *http.Request) string {
	return strings.ToLower(r.Header.Get("Accept"))
}

func marshal(w http.ResponseWriter, r *http.Request, reply *spec.IPReply) (body []byte, err error) {
	accept := getAcceptHeader(r)
	w.Header().Set("Content-Type", accept)
	switch accept {
	case "application/protobuf":
		return proto.Marshal(reply)
	case "application/json":
		return json.Marshal(reply)
	default:
		w.Header().Set("Content-Type", "text/plain")
		return []byte(reply.Ip), nil
	}
}

// Handler responds with the IP address of the request.
func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := marshal(w, r, resolver(r))
	respond(w, r, body, err)
}
