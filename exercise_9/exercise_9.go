package exercise_9

import "fmt"

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

func one(in chan int) {
	// Пишем данные в канал
	for i := 0; i < 10; i++ {
		in <- i
	}

	close(in) // Закрываем канал, т.к. он больше не будет нужен
}

func two(in, out chan int) {
	defer close(out) // Закроем канал после return

	// В бесконечном цикле, чтобы принять все данные
	for {
		// Пишем из in в num до тех пор, пока num открыт
		// если !open - завершаем выполнение блока for
		num, open := <-in
		if !open {
			return
		}

		// Пишем в канал x*2
		out <- num * 2
	}
}

func Run() {
	var in = make(chan int)
	var out = make(chan int)

	go one(in)
	go two(in, out)

	// Выводим результат
	for i := range out {
		fmt.Println(i)
	}
}
