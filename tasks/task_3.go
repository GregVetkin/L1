package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)



func SqrSum(numbers []int64) int64{
	var sum int64
	// wait group для ожидания выполнения всех рабочих горутин
	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for _, number := range numbers{
		go func(number int64) {
			defer wg.Done()
			// используем атомик для безопасной работы со счетчиком суммы
			atomic.AddInt64(&sum, number * number)
		}(number)
	}
	wg.Wait()
	return sum
}



func main() {
	var numbers = []int64{2, 4, 6, 8, 10}


	fmt.Println(SqrSum(numbers))	// 220


}