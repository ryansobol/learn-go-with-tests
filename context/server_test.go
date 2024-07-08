package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type StoreSpy struct {
	response string
	t        *testing.T
}

func (s *StoreSpy) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string

		for _, char := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(char)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type ResponseWriterSpy struct {
	written bool
}

func (r *ResponseWriterSpy) Header() http.Header {
	r.written = true
	return nil
}

func (r *ResponseWriterSpy) Write([]byte) (int, error) {
	r.written = true
	return 0, errors.New("not implemented")
}

func (r *ResponseWriterSpy) WriteHeader(statusCode int) {
	r.written = true
}

func TestServer(t *testing.T) {
	t.Run("writes data to store", func(t *testing.T) {
		expected := "Hello, Jane!"
		store := &StoreSpy{response: expected, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		actual := response.Body.String()

		assert.Equal(t, expected, actual)
	})

	t.Run("cancel store's work when request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &StoreSpy{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())

		// cancel()

		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)

		response := &ResponseWriterSpy{}

		server.ServeHTTP(response, request)

		assert.False(t, response.written)
	})
}
