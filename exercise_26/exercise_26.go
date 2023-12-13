package exercise_26

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func isUnique(s string) bool {
	// Так как Go использует UTF-8 для строк, обращаясь по индексу
	// элемента в строке вместо буквы мы получим байтовое представление
	// символа - руну. Поэтому мапа из ключей rune и значений типа bool
	m := make(map[rune]bool)

	// приводим строку к нижнему регистру
	s = strings.ToLower(s)

	for _, char := range s {
		// Если руна уже есть в мапе, тогда возвращаем false
		if m[char] {
			return false
		}
		// иначе seen[char] = true
		m[char] = true
	}

	return true
}

func Run() {
	values := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, v := range values {
		fmt.Println(isUnique(v))
	}
}
