package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func WorkerPrinter(ctx context.Context, inChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var number int
	for {
		select {
		// Выводим в консоль число, если канал не пустой
		case number = <-inChan:
			fmt.Println(number)
		// Если получаем от контекста сигнал о завершении - выходим из фунции
		case <-ctx.Done():
			fmt.Println("Closing goroutine...")
			return
		}
	}
}


func main() {
	var wg sync.WaitGroup
	var NumberOfWorkers = 5

	wg.Add(NumberOfWorkers)
	// Создаем канал с буфером размера количества воркеров
	inChan := make(chan int, NumberOfWorkers)
	defer close(inChan)

	ctx, cancel := context.WithCancel(context.Background())


	// Запускаем воркеры
	for i := 0; i < NumberOfWorkers; i++ {
		go WorkerPrinter(ctx, inChan, &wg)
	}

	// Запускаем воркер, который записывает в канал числа
	go func() {
		// Добавляем в группу ожидания 1 (эта функция)
		wg.Add(1)
		defer wg.Done()
		for {
			select{
			// Если придет сигнал о завершении - воркер записи завершит работы
			case <-ctx.Done():
				fmt.Println("Writer stopped...")
				return
			// Без сигнала завершения - будет срабаывать default case записи случайного числа в канал
			default:
				inChan <- rand.Intn(101)
			}
		}
	}()


	// создаем канал передачи сигнала от ос
	stopSignalCh := make(chan os.Signal)
	// Добавляем уведомление при сигнале прерывания от ос
	signal.Notify(stopSignalCh, os.Interrupt)

	// Ждем в главной горутине сигнала
	<-stopSignalCh
	fmt.Println("Stopping process...")
	cancel()
	wg.Wait()
	fmt.Println("The process stopped gracefully")
}
