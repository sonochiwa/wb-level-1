package exercise_19

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func Run() {
	word := "главрыба"
	runes := []rune(word)
	wordLen := len(runes)

	// Переворачиваем слово
	for i := 0; i < wordLen/2; i++ {
		runes[i], runes[wordLen-i-1] = runes[wordLen-i-1], runes[i]
	}

	fmt.Println(string(runes))
}
