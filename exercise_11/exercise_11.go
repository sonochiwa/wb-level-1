package exercise_11

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func intersection(set1, set2 []int) []int {
	var result []int
	var inUse map[int]bool

	// Заполняем карту элементами из первого множества
	for _, val := range set1 {
		inUse[val] = true
	}

	// Проверяем каждый элемент второго множества
	// и добавляем его в результат, если он есть в первом множестве
	for _, val := range set2 {
		if inUse[val] {
			result = append(result, val)
		}
	}

	return result
}

func Run() {
	// Множества
	set1 := []int{2, 1, 3, 5, 4}
	set2 := []int{7, 4, 5, 6, 3}

	result := intersection(set1, set2)

	fmt.Println("Пересечение:", result)
}
