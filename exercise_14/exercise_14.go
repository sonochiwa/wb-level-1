package exercise_14

import (
	"fmt"
)

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

func checkType(value interface{}) {

	// type switch конструкция
	switch value.(type) {
	case int:
		fmt.Println("Value is int")
	case string:
		fmt.Println("Value is string")
	case bool:
		fmt.Println("Value is bool")
	case chan any:
		fmt.Println("Value is chan")
	}
}

func Run() {
	var x interface{} = true
	checkType(x)
}
