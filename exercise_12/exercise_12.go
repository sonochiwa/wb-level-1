package exercise_12

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.

func set(words []string) map[string]struct{} {
	set := make(map[string]struct{})

	// Преобразуем слайс в множество
	for _, word := range words {
		set[word] = struct{}{}
	}

	return set
}

func Run() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	result := set(words)

	fmt.Println(result)
}
