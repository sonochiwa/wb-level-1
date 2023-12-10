package exercise_2

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

// Функция подсчета квадрата числа.
// Принимаем wg по указателю, чтобы можно было менять значение счетчика, объявленного в другой области видимости.
func square(value int, wg *sync.WaitGroup, ch chan int) {
	// Декрементим счетчик wg.
	defer wg.Done()

	// Считаем квадрат переданного в функцию числа.
	square := value * value

	// Отправляем данные в канал.
	ch <- square
}

func Run() {
	// Объявляем счетчик wg.
	var wg sync.WaitGroup

	// Инициализируем массив значений типа int.
	numbers := [5]int{2, 4, 6, 8, 10}

	// Инициализируем канал для значений типа int c размером буфера равным длине массива numbers.
	ch := make(chan int, len(numbers))

	// Пробегаемся по элементам массива numbers.
	for _, num := range numbers {
		// Инкрементим счетчик wg для каждой запущенной горутины.
		wg.Add(1)

		// Считаем квадрат числа
		go square(num, &wg, ch)
	}

	// Блокируем основную горутину, пока счетчик wg снова не станет равным нулю.
	// Это будет означать, что все горутины выполнились.
	wg.Wait()

	// Закрываем канал.
	close(ch)

	// Выводим значения квадратов из буфера канала ch
	for v := range ch {
		fmt.Println(v)
	}
}
