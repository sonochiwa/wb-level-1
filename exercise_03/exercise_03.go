package exercise_03

import "fmt"

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2 + 3^2 + 4^2 ...) с использованием конкурентных вычислений.

func square(value int, ch chan int) {
	square := value * value

	// Записываем значение в канал
	ch <- square
}

func Run() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	sum := 0

	for _, num := range numbers {
		go square(num, ch)
		sum += <-ch
	}

	close(ch)
	fmt.Println(sum)
}
