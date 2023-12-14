package exercise_8

import (
	"fmt"
	"strconv"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в
// 1 или 0.

func replaceBit(value int64, index int, replacement int) (int64, error) {
	// Переводим из десятичной системы в двоичную
	int64ToBit := strconv.FormatInt(value, 2)

	// Меняем i-й бит и собираем в строку
	bitToStr := int64ToBit[:index] + strconv.Itoa(replacement) + int64ToBit[index+1:]

	// Строку парсим из двоичной системы в десятичную
	return strconv.ParseInt(bitToStr, 2, 64)
}

func Run() {
	value, _ := replaceBit(23, 2, 0)
	fmt.Println(value)
}
