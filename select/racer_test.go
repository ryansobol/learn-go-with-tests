package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRacer(t *testing.T) {
	t.Run("return the fastest URL", func(t *testing.T) {
		slowServer := makeHttpServerWithDelay(20 * time.Millisecond)
		fastServer := makeHttpServerWithDelay(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expected := fastURL
		actual, err := Racer(slowURL, fastURL)

		assert.NoError(t, err)

		assert.Equal(t, expected, actual)
	})

	t.Run("return an error on timeout", func(t *testing.T) {
		server := makeHttpServerWithDelay(25 * time.Millisecond)

		defer server.Close()

		url := server.URL

		_, err := RacerConfigurable(url, url, 10*time.Millisecond)

		assert.Error(t, err)
	})
}

func makeHttpServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}),
	)
}
