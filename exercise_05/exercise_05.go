package exercise_05

import (
	"fmt"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

// Функция для отправки значений в бесконечном цикле
func sendData(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; ; i++ {
		// Отправляем значение
		ch <- i

		// Ждем 1 секунду
		time.Sleep(time.Second)
	}
}

// Функция для чтения данных из канала
func getData(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func Run() {
	ch := make(chan int)
	var wg sync.WaitGroup

	// Запускаем горутину для отправки значений в канал
	wg.Add(1)

	// Отправляем данные в канал
	go sendData(ch, &wg)

	// Читаем данные из канала
	go getData(ch)

	// Ждем N секунд
	time.Sleep(time.Second * time.Duration(3))

	// Сигнализируем горутине завершиться, т.к. она выполняется в бесконечном цикле
	// и сама не завершится
	wg.Done()

	// Ждем завершения горутины
	wg.Wait()
}
