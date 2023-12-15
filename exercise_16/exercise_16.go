package exercise_16

import (
	"fmt"
	"sort"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка.

// Функция сортировки
func quickSort(a *[]int) {
	sort.Slice(*a, func(i, j int) bool {
		return (*a)[i] < (*a)[j]
	})
}

func Run() {
	a := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	quickSort(&a)
	fmt.Println(a) // [2 3 5 6 7 9 10 11 12 14]
}
