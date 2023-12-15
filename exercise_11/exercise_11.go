package exercise_11

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func intersection(set1, set2 map[int]struct{}) []int {
	var result []int
	inUse := make(map[int]bool)

	// Заполняем карту элементами из первого множества
	for k, _ := range set1 {
		inUse[k] = true
	}

	// Проверяем каждый элемент второго множества
	// и добавляем его в результат, если он есть в первом множестве
	for k, _ := range set2 {
		if inUse[k] {
			result = append(result, k)
		}
	}

	return result
}

func Run() {
	// Инициализируем множества с дефолтными значениями
	set1 := map[int]struct{}{1: {}, 2: {}, 3: {}}
	set2 := map[int]struct{}{2: {}, 3: {}, 4: {}}

	result := intersection(set1, set2)

	fmt.Println("Пересечение:", result)
}
