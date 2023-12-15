package exercise_17

import (
	"fmt"
	"sort"
)

// Реализовать бинарный поиск встроенными методами языка.

func binarySearch(a []int, x int) int {
	index := sort.SearchInts(a, x)

	if a[index] == x {
		return index
	}

	return -1
}

func Run() {
	a := []int{1, 2, 3, 4}
	x := binarySearch(a, 2)

	fmt.Println("Индекс искомого значения в наборе данных:", x)
}
