package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls += 1
}

const (
	sleep = "sleep"
	write = "write"
)

type OperationsSpy struct {
	Calls []string
}

func (o *OperationsSpy) Sleep() {
	o.Calls = append(o.Calls, sleep)
}

func (o *OperationsSpy) Write(p []byte) (n int, err error) {
	o.Calls = append(o.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("writes 3 to Go!", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spy := SleeperSpy{}

		Countdown(&buffer, &spy)

		actual := buffer.String()
		expected := `3
2
1
Go!`

		assert.Equal(t, expected, actual)

		assert.Equal(t, 3, spy.Calls)
	})

	t.Run("sleep before every write", func(t *testing.T) {
		spy := OperationsSpy{}

		Countdown(&spy, &spy)

		expected := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		assert.Equal(t, expected, spy.Calls)
	})
}

type TimeSpy struct {
	durationSlept time.Duration
}

func (t *TimeSpy) Sleep(duration time.Duration) {
	t.durationSlept = duration
}

func TestConfig(t *testing.T) {
	sleepTime := 5 * time.Second
	spy := TimeSpy{}
	sleeper := SleeperConfig{sleepTime, spy.Sleep}

	sleeper.Sleep()

	assert.Equal(t, spy.durationSlept, sleepTime)
}
