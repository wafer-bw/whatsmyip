package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/wafer-bw/whatsmyip/spec"
	"google.golang.org/protobuf/proto"
)

// Handler is an HTTP handler that responds with the IP address of the request.
//
// This is named Handler to satisfy vercel's serverless function requirements.
func Handler(w http.ResponseWriter, r *http.Request) {
	reply, err := identifyIP(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	accept := strings.ToLower(r.Header.Get("accept"))

	contentType, body, err := marshalIP(accept, reply)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", contentType)
	_, _ = w.Write(body)
}

// identifyIP address of the request.
func identifyIP(r *http.Request) (*spec.IPReply, error) {
	ip := r.Header.Get("x-real-ip")
	if ip == "" {
		ip = r.Header.Get("x-forwarded-for")
	}
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	if ip == "" {
		return nil, ErrNoIP
	}

	return &spec.IPReply{Ip: ip}, nil
}

// marshalIP reply to bytes and a content-type string based on provided accept string.
//
// If accept is blank then reply.Ip is returned as text/plain.
func marshalIP(accept string, reply *spec.IPReply) (string, []byte, error) {
	switch accept {
	case "application/protobuf":
		b, err := proto.Marshal(reply)
		return accept, b, err
	case "application/json":
		b, err := json.Marshal(reply)
		return accept, b, err
	default:
		return "text/plain", []byte(reply.Ip), nil
	}
}

// Error type permits sentinel errors defined as constants.
type Error string

func (e Error) Error() string { return string(e) }

// ErrNoIP occurs in the unlikely event that the request IP address could not
// be identified.
const ErrNoIP = Error("unable to identify IP address")
