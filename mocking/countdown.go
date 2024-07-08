package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SleeperConfig struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *SleeperConfig) Sleep() {
	s.sleep(s.duration)
}

const (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

// type SleeperDefault struct{}

// func (s SleeperDefault) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

func main() {
	sleeper := SleeperConfig{1 * time.Second, time.Sleep}

	Countdown(os.Stdout, &sleeper)
}
