package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Ожидание закрытия канала в range
func GoroutineChannelClose_1(ch <- chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range ch {
		fmt.Println(n)
	}
	fmt.Println("Channel closed in range loop, ending goroutine")
}

// Ожидание закрытия канала явно
func GoroutineChannelClose_2(ch <- chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		n, ok := <- ch
		if !ok {
			fmt.Println("Channel closed in for cycle, ending goroutine")
			return
		}
		fmt.Println(n)
	}
}


// Ожидание закрытия канала явно
func GoroutineClosingChanel(ch <- chan int, closeCh <- chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case n := <-ch:
			fmt.Println(n)
		case <-closeCh:
			fmt.Println("Signal from closeCh, ending goroutine")
			return
		}
	}
}

// Завершение по таймеру
func GoroutineClosingByTimer(ch <- chan int, t *time.Timer, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case n := <-ch:
			fmt.Println(n)
		case <-t.C:
			fmt.Println("Signal from timer, ending goroutine")
			return
		}
	}
}

// Завершение по сигналу из контекста
func GoroutineClosingByContextSignal(ch <- chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case n := <-ch:
			fmt.Println(n)
		case <-ctx.Done():
			fmt.Println("Signal from context, ending goroutine")
			return
		}
	}
}



func main() {
	var wg sync.WaitGroup

	// ============================
	// Закрытие канала 1
	ch := make(chan int)
	wg.Add(1)
	go GoroutineChannelClose_1(ch, &wg)
	ch <- 101
	ch <- 102
	time.Sleep(1 * time.Second)
	close(ch)
	wg.Wait()
	// ============================

	// ============================
	// Закрытие канала 2
	ch = make(chan int)
	wg.Add(1)
	go GoroutineChannelClose_2(ch, &wg)
	ch <- 201
	ch <- 202
	time.Sleep(1 * time.Second)
	close(ch)
	wg.Wait()
	// ============================

	// ============================
	// Сигнал в отдельный канал о закрытии
	ch = make(chan int)
	closeCh := make(chan bool)
	wg.Add(1)
	go GoroutineClosingChanel(ch, closeCh, &wg)
	ch <- 301
	ch <- 302
	time.Sleep(1 * time.Second)
	closeCh <- true
	wg.Wait()
	close(ch)
	close(closeCh)
	// ============================

	// Сигнал в отдельный канал о закрытии
	ch = make(chan int)
	timer := time.NewTimer(1 * time.Second)
	wg.Add(1)
	go GoroutineClosingByTimer(ch, timer, &wg)
	ch <- 401
	ch <- 402
	wg.Wait()
	timer.Stop()
	close(ch)
	// ============================	

	// Сигнал из контекста
	ch = make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go GoroutineClosingByContextSignal(ch, ctx, &wg)
	ch <- 501
	ch <- 502
	time.Sleep(1 * time.Second)
	cancel()
	wg.Wait()
	close(ch)
	// ============================	
}