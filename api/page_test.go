package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/whatsmyip/api"
)

func TestPage(t *testing.T) {
	t.Run("respond with the homepage with the IP address already in it", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()
		http.HandlerFunc(api.Page).ServeHTTP(recorder, request)
		resp := recorder.Result()
		defer func() { _ = resp.Body.Close() }()
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, "text/html; charset=utf-8", resp.Header.Get("content-type"))
		require.Contains(t, string(body), `<div id="ip">192.0.2.1</div>`)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("escape the IP address header", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set("X-Real-Ip", "<script>")
		recorder := httptest.NewRecorder()
		http.HandlerFunc(api.Page).ServeHTTP(recorder, request)
		resp := recorder.Result()
		defer func() { _ = resp.Body.Close() }()
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Contains(t, string(body), `<div id="ip">&lt;script&gt;</div>`)
	})

	t.Run("respond with an error when the IP address can't be identified", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.RemoteAddr = ""
		recorder := httptest.NewRecorder()
		http.HandlerFunc(api.Page).ServeHTTP(recorder, request)
		require.Equal(t, http.StatusInternalServerError, recorder.Code)
		require.Equal(t, api.ErrNoIP.Error()+"\n", recorder.Body.String())
	})
}
