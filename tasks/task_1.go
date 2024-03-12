package main

import "fmt"

type Human struct {
	Name  string
	Age   int
	Title string
}

func (h *Human) SetName(name string) {
	h.Name = name
}

func (h *Human) SetTitle(title string) {
	h.Title = title
}

func (h *Human) Birthday() {
	h.Age++
}

func (h *Human) Info() {
	fmt.Printf("Human with name %s, age %d and title '%s'\n", h.Name, h.Age, h.Title)
}

// ==========================================================

type Action struct {
	Title string
	Human
}

func (a *Action) SetTitle(title string) {
	a.Title = title
}

func main() {
	var TestHuman = Human{
		Name:  "Charles I",
		Age:   48,
		Title: "King",
	}

	var TestAction = Action{
		Title: "Declaration of war",
		Human: TestHuman,
	}

	// Поле не уникально, след-но выведется поле более верхнего уровня, те структуры Action
	fmt.Println(TestAction.Title) // Declaration of war

	// К уникальным полям можно обращаться сразу из верхней структуры
	fmt.Println(TestAction.Age) // 48
	// Ровно как и к методам
	TestAction.Birthday()
	fmt.Println(TestAction.Age) // 49

	// Если методы не уникальны, то вызывается метод выше по уровню
	fmt.Println(TestAction.Title) // Declaration of war
	TestAction.SetTitle("Sleeping")
	// Сменилось поле Title у Action но не у Human
	fmt.Println(TestAction.Title)       // Sleeping
	fmt.Println(TestAction.Human.Title) // King

	//Для вызова метода SetTitle структуры Human (не уникальной), нужно явно обраться к структуре Human
	TestAction.Human.SetTitle("Peasant")
	fmt.Println(TestAction.Human.Title) // Peasant

	fmt.Println(TestAction)

}
