package exercise_20

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func Reverse(phrase *string) {
	// Преобразовываем фразу в слайс строк
	words := strings.Split(*phrase, " ")
	length := len(words)

	// Разворачиваем слайс строк
	for i := 0; i < length/2; i++ {
		words[i], words[length-i-1] = words[length-i-1], words[i]
	}

	// Склеиваем слайс строк в строку
	*phrase = strings.Join(words, " ")
}

func Run() {
	phrase := "snow dog sun trash"
	Reverse(&phrase)
	fmt.Println(phrase)
}
