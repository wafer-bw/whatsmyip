package api_test

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/whatsmyip/api"
)

var router *mux.Router
var url = "http://localhost:1234/"

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)
	router = api.GetRouter()
	os.Exit(m.Run())
}

func mockRequest(method string, url string, headers map[string]string, bodyReader io.Reader) ([]byte, *http.Response, error) {
	request := httptest.NewRequest(method, url, bodyReader)
	for key, val := range headers {
		request.Header.Set(key, val)
	}
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	resp := recorder.Result()
	body, err := io.ReadAll(recorder.Body)
	return body, resp, err
}

func TestHandlerOkText(t *testing.T) {
	body, resp, err := mockRequest(http.MethodGet, url, nil, nil)
	defer resp.Body.Close()
	require.NoError(t, err)
	require.Equal(t, "text/plain", resp.Header.Get("content-type"))
	require.Equal(t, "192.0.2.1", string(body))
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHandlerOkJson(t *testing.T) {
	requestHeaders := map[string]string{"Accept": "application/json"}
	body, resp, err := mockRequest(http.MethodGet, url, requestHeaders, nil)
	defer resp.Body.Close()
	require.NoError(t, err)
	require.Equal(t, "application/json", resp.Header.Get("content-type"))
	require.Equal(t, "{\"ip\":\"192.0.2.1\"}", string(body))
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHandlerOkProto(t *testing.T) {
	requestHeaders := map[string]string{"Accept": "application/protobuf"}
	body, resp, err := mockRequest(http.MethodGet, url, requestHeaders, nil)
	defer resp.Body.Close()
	require.NoError(t, err)
	require.Equal(t, "application/protobuf", resp.Header.Get("content-type"))
	require.Equal(t, "\n\t192.0.2.1", string(body))
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
