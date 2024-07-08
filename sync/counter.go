package main

import "sync"

type Counter struct {
	mutex sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value += 1
}

func (c *Counter) Value() int {
	return c.value
}
