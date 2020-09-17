package main

import (
	"io"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var url = "http://localhost:1234/"

func mockRequest(method string, url string, headers map[string]string, body io.Reader) ([]byte, int, error) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", url, nil)
	for key, val := range headers {
		request.Header.Set(key, val)
	}
	router := getRouter()
	router.ServeHTTP(recorder, request)
	result, err := ioutil.ReadAll(recorder.Body)
	return result, recorder.Result().StatusCode, err
}

func TestHandlerOk(t *testing.T) {
	body, status, err := mockRequest("GET", url, nil, nil)
	require.NoError(t, err)
	require.Equal(t, "192.0.2.1", string(body))
	require.Equal(t, 200, status)
}

func TestHandlerOkJson(t *testing.T) {
	headers := map[string]string{"Accept": "application/json"}
	body, status, err := mockRequest("GET", url, headers, nil)
	require.NoError(t, err)
	require.Equal(t, "{\"ip\":\"192.0.2.1\"}", string(body))
	require.Equal(t, 200, status)
}

func TestHandlerOkProto(t *testing.T) {
	headers := map[string]string{"Accept": "application/protobuf"}
	body, status, err := mockRequest("GET", url, headers, nil)
	require.NoError(t, err)
	require.Equal(t, "\n\t192.0.2.1", string(body))
	require.Equal(t, 200, status)
}

func TestGetEnv(t *testing.T) {
	v := getEnv("WMIPTESTENV", "DEFAULT")
	require.Equal(t, "DEFAULT", v)
	os.Setenv("WMIPTESTENV", "NONDEFAULT")
	v = getEnv("WMIPTESTENV", "DEFAULT")
	require.Equal(t, "NONDEFAULT", v)
}
