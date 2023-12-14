package exercise_22

import (
	"fmt"
	"math"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

// Функция для умножения
func multiply(a, b float64) float64 {
	return a * b
}

// Функция для деления
func divide(a, b float64) float64 {
	return a / b
}

// Функция для сложения
func sum(a, b float64) float64 {
	return a + b
}

// Функция для вычитания
func subtract(a, b float64) float64 {
	return a - b
}

func Run() {
	// Присвоим значения переменным a и b, которые больше чем 2^20
	a := math.Pow(2, 21)
	b := math.Pow(2, 22)

	resultMultiply := multiply(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, resultMultiply)

	resultDiv := divide(a, b)
	fmt.Printf("%v / %v = %v\n", a, b, resultDiv)

	resultAdd := sum(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, resultAdd)

	resultSub := subtract(a, b)
	fmt.Printf("%v - %v = %v\n", a, b, resultSub)
}
