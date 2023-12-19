package exercise_04

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из
// канала и выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.

// Получаем количество воркеров
func getWorkersCount() int {
	// Читаем данные из консоли
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите количество воркеров: ")
	scanner.Scan()

	// Конвертируем строку в число
	workersCount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Printf("\"%s\" некорректное значение. \nОшибка: %s\n", scanner.Text(), err)
		return 0
	}

	return workersCount
}

// Создаем воркер
func addWorker(i int, ch chan int, signInterrupt <-chan os.Signal) {
	for {
		select {
		case <-ch:
			// Отправляем воркер в ожидание
			time.Sleep(time.Duration(<-ch) * time.Millisecond)

			// Передаем данные в канал
			fmt.Printf("[Worker %d] Sleep %d milliseconds\n", i, <-ch)
		case <-signInterrupt:
			fmt.Printf("Worker %d finished.\n", i)

			// Завершаем функцию тем самым выключая горутину
			return
		}
	}
}

func Run() {
	workersCount := getWorkersCount()

	ch := make(chan int, workersCount)
	signInterrupt := make(chan os.Signal)

	// Отправим в канал сигнал по нажатию CTRL+C
	signal.Notify(signInterrupt, os.Interrupt)

	// Запускаем указанное количество воркеров
	for i := 1; i <= workersCount; i++ {
		go addWorker(i, ch, signInterrupt)
	}

	for {
		select {
		// На каждой итерации цикла отправляем случайное число в канал
		case ch <- rand.Intn(3000):

		// Останавливаем горутины
		case <-signInterrupt:
			log.Println("Shutdown goroutines")

			// Закрываем канал signInterrupt сообщая всем воркерам о завершении работы
			close(signInterrupt)

			// Завершаем программу
			os.Exit(0)
		}
	}
}
