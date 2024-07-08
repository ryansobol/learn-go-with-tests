package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter 3 times serially", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		actual := counter.Value()
		expected := 3

		assert.Equal(t, expected, actual)
	})

	t.Run("increme counter 1,000 times concurrently", func(t *testing.T) {
		expected := 1000
		counter := Counter{}

		var wg sync.WaitGroup

		wg.Add(expected)

		for i := 0; i < expected; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		actual := counter.Value()

		assert.Equal(t, expected, actual)
	})
}
