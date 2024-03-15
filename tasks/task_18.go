package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type CounterAtomic struct {
	value int64
}

func (c *CounterAtomic) Add() {
	atomic.AddInt64(&c.value, 1)
}

func (c CounterAtomic) String() string {
	return fmt.Sprintf("Counter-Atomic: %v", c.value)
}


type CounterMutex struct {
	value int64
	sync.Mutex
}

func (c *CounterMutex) Add() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.value++
}

func (c CounterMutex) String() string {
	return fmt.Sprintf("Counter-Mutex: %v", c.value)
}


func main() {
	counterAtomic := CounterAtomic{}
    for i := 0; i < 100; i++ {
        go func(no int) {
			for i := 0; i < 10000; i++ {
                counterAtomic.Add()
            }
        }(i)
    }
    time.Sleep(time.Second)
    fmt.Println(counterAtomic) // 1000000


	counterMutex := CounterMutex{}
    for i := 0; i < 100; i++ {
        go func(no int) {
			for i := 0; i < 10000; i++ {
                counterMutex.Add()
            }
        }(i)
    }
    time.Sleep(time.Second)
    fmt.Println(counterMutex) // 1000000
}