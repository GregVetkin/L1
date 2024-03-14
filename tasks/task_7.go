package main

import (
	"fmt"
	"sync"
)

type ConcurMap[T any] struct {
	sync.RWMutex
	data	map[string]T
}


func NewConcurMap[T any]() ConcurMap[T] {
	return ConcurMap[T]{data: make(map[string]T)}
}


func (c *ConcurMap[T]) Get(key string) (T, bool) {
	c.RLock()
	defer c.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

func (c *ConcurMap[T]) Set(key string, value T) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}



func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	concmap := NewConcurMap[string]()

	concmap.Set("Russia", "Moscow")
	concmap.Set("Egypt", "Cairo")

	go func() {
		defer wg.Done()
		fmt.Println(concmap.Get("Russia"))
		fmt.Println(concmap.Get("Spain"))
	}()

	go func() {
		defer wg.Done()
		concmap.Set("Spain", "Madrid")
		fmt.Println(concmap.Get("Egypt"))
	}()


	wg.Wait()
}