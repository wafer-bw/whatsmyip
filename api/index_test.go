package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/whatsmyip/api"
)

func TestIP(t *testing.T) {
	t.Run("respond with IP address as text/plain content-type by default", func(t *testing.T) {
		body, resp, err := mockRequest(nil, nil)
		defer resp.Body.Close()
		require.NoError(t, err)
		require.Equal(t, "text/plain", resp.Header.Get("content-type"))
		require.Equal(t, "192.0.2.1", string(body))
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("respond with IP address as application/json content-type via accept header", func(t *testing.T) {
		requestHeaders := map[string]string{"Accept": "application/json"}
		body, resp, err := mockRequest(requestHeaders, nil)
		defer resp.Body.Close()
		require.NoError(t, err)
		require.Equal(t, "application/json", resp.Header.Get("content-type"))
		require.Equal(t, `{"ip":"192.0.2.1"}`, string(body))
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("respond with IP address as application/protobuf content-type via accept header", func(t *testing.T) {
		requestHeaders := map[string]string{"Accept": "application/protobuf"}
		body, resp, err := mockRequest(requestHeaders, nil)
		defer resp.Body.Close()
		require.NoError(t, err)
		require.Equal(t, "application/protobuf", resp.Header.Get("content-type"))
		require.Equal(t, "\n\t192.0.2.1", string(body))
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func mockRequest(headers map[string]string, bodyReader io.Reader) ([]byte, *http.Response, error) {
	request := httptest.NewRequest(http.MethodGet, "/", bodyReader)
	for key, val := range headers {
		request.Header.Set(key, val)
	}
	recorder := httptest.NewRecorder()
	router := http.HandlerFunc(api.Handler)
	router.ServeHTTP(recorder, request)
	resp := recorder.Result()
	body, err := io.ReadAll(recorder.Body)
	return body, resp, err
}
