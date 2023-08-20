package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wafer-bw/whatsmyip/api"
)

func BenchmarkHandler(b *testing.B) {
	handler := http.HandlerFunc(api.Handler)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	b.Run("text", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			handler.ServeHTTP(w, r)
		}
	})

	r.Header.Set("Accept", "application/protobuf")
	b.Run("protobuf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			handler.ServeHTTP(w, r)
		}
	})

	r.Header.Set("Accept", "application/json")
	b.Run("json", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			handler.ServeHTTP(w, r)
		}
	})
}
