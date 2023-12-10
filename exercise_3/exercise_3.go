package exercise_3

import (
	"fmt"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func square(value int, ch chan int) {
	square := value * value

	// Записываем значение в канал
	ch <- square
}

func Run() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	// Сумма квадратов числел
	sum := 0

	for _, num := range numbers {
		go square(num, ch)

		// Увеличиваем значение sum на квадрат числа из канала
		sum += <-ch
	}

	close(ch)
	fmt.Println(sum)
}
