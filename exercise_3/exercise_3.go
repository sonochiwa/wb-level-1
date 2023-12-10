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

	for _, num := range numbers {
		go square(num, ch)
		fmt.Println(<-ch)
	}

	close(ch)
}
