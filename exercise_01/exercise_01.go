package exercise_01

import "fmt"

//	Дана структура human (с произвольным набором полей и методов).
//	Реализовать встраивание методов в структуре action от родительской
//	структуры human (аналог наследования).

// Структура human с произвольными полями и методами.
type human struct {
	name string
	age  int
}

// Структура action встраивает структуру human.
type action struct {
	human // встраиваем human в action
}

// Метод greeting для структуры human.
func (h *human) greeting() {
	fmt.Printf("Hi, my name is %s. I am %d years old.\n", h.name, h.age)
}

// Метод sleep для структуры action.
func (a *action) sleep() {
	fmt.Println(a.human.name, "is sleeping.")
}

func Run() {
	// Создаем объект person
	person := action{
		human: human{
			name: "Rashid",
			age:  23,
		},
	}

	// Методы human доступны через встраивание.
	person.greeting()
	person.sleep()
}
