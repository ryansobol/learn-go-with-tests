package main

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(a, b string) (winner string) {
// 	durationA := measureResponseTime(a)
// 	durationB := measureResponseTime(b)

// 	if durationA < durationB {
// 		return a
// 	}

// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	startTime := time.Now()
// 	http.Get(url)
// 	return time.Since(startTime)
// }

var timeoutDefault = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return RacerConfigurable(a, b, timeoutDefault)
}

func RacerConfigurable(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
