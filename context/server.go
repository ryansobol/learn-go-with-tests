package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		data, err := store.Fetch(req.Context())
		if err != nil {
			return
		}

		fmt.Fprint(writer, data)
	}
}
