package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	//mu sync.Mutex
	mu sync.RWMutex
	c  map[string]int
}

func (c *Counter) CountMe() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c

}

func (c *Counter) CountMeAgain() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c

}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	c.c[key]++
	c.mu.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	val := c.c[key]
	if val == 30 {
		return c.c[key] + 30
	}
	return c.c[key]
	//defer c.mu.Unlock()
	//return c.c[key]
}

func main() {
	key := "test"

	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}

	time.Sleep(time.Second * 1)
	fmt.Println(c.Value(key))
}
