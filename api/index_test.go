package api

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var url = "http://localhost:1234/"

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func mockRequest(method string, url string, headers map[string]string, body io.Reader) ([]byte, http.Header, int, error) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", url, nil)
	for key, val := range headers {
		request.Header.Set(key, val)
	}
	router := GetRouter()
	router.ServeHTTP(recorder, request)
	result, err := ioutil.ReadAll(recorder.Body)
	return result, recorder.Result().Header, recorder.Result().StatusCode, err
}

func TestHandlerOkText(t *testing.T) {
	body, headers, status, err := mockRequest("GET", url, nil, nil)
	require.NoError(t, err)
	require.Equal(t, "text/plain", headers.Get("content-type"))
	require.Equal(t, "192.0.2.1", string(body))
	require.Equal(t, 200, status)
}

func TestHandlerOkJson(t *testing.T) {
	requestHeaders := map[string]string{"Accept": "application/json"}
	body, headers, status, err := mockRequest("GET", url, requestHeaders, nil)
	require.NoError(t, err)
	require.Equal(t, "application/json", headers.Get("content-type"))
	require.Equal(t, "{\"ip\":\"192.0.2.1\"}", string(body))
	require.Equal(t, 200, status)
}

func TestHandlerOkProto(t *testing.T) {
	requestHeaders := map[string]string{"Accept": "application/protobuf"}
	body, headers, status, err := mockRequest("GET", url, requestHeaders, nil)
	require.NoError(t, err)
	require.Equal(t, "application/protobuf", headers.Get("content-type"))
	require.Equal(t, "\n\t192.0.2.1", string(body))
	require.Equal(t, 200, status)
}
