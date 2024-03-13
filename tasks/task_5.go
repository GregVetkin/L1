package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var lifetime = 1 * time.Second

	ch := make(chan int)

	timer := time.NewTimer(lifetime)
	defer timer.Stop()

	// горутина записи в канал
	go func() {
		defer wg.Done()
		for {
			select {
			case <-timer.C:
				fmt.Println("Writer: Time out")
				close(ch)
				return
			default:
				ch <- rand.Intn(11)
				fmt.Println("Writer: Data sent")
			}
		}
	}()

	// горутина чтения и зканала, пока он не окажется закрыт
	go func() {
		var n int
		defer wg.Done()
		for n = range ch {
			fmt.Println("Reader: Data received ", n)
		}
		fmt.Println("Reader: Channel closed")
	}()


	wg.Wait()
	fmt.Println("End")
}
