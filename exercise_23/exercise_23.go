package exercise_23

import "fmt"

// Удалить i-ый элемент из слайса.

// Удаляем элемент в исходном массиве по указателю.
func removeElement(data *[]int, index int) {
	// Нативно поддержки remove из слайса нет, поэтому
	// заполняем исходный слайс его срезами до и после index.
	*data = append((*data)[:index], (*data)[index+1:]...)
}

func Run() {
	data := []int{1, 2, 3, 4, 5}
	removeElement(&data, 0)
	fmt.Println(data)
}
