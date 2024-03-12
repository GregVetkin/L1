package main

import (
	"fmt"
	"sync"
)

func Squares(numbers []int) {
	var wg sync.WaitGroup

	// Добавляем ожидаемое число "работ" в счетчик ожидания
	wg.Add(len(numbers))

	for _, number := range numbers {
		go square(number, &wg)
	}

	// Ождием выполнения всех задач, чтобы освободить данную горутину
	wg.Wait()
}


func square(number int, wg *sync.WaitGroup) {
	// После отработки функции сообщаем группе ожидания, что работа у данной горутины сделана
	defer wg.Done()
	fmt.Println(number * number)
}


func main() {
	var numbers = []int{2, 4, 6, 8, 10}
	
	/* 
		Горутины посчитают квадрат числа в слайсе и напечатают результат
	 	WaitGroup создается чтобы основная горутина не завершилась раньше "наших" рабочих
	 	Горутины получают доступ к разным элементам слайса, 
	 	но какая из горутин быстрее выполнит задачу - не детерменированно.
		Отсюда при каждом запуске результат вывода будет в случайном порядке
	*/
	Squares(numbers)
}
