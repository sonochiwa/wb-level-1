package exercise_6

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// 1. Использование каналов для сигнализации завершения
func methodOne() {
	// По истечении времени отправляем в канал stop данные
	// из-за чего в операторе select { ... } срабатывает блок stop, в котором
	// мы завершаем горутину вызовом return

	stop := make(chan bool)

	go func() {
		for {
			select {
			default:
				fmt.Println("Goroutine is running...")
				time.Sleep(time.Second)
			case <-stop:
				fmt.Println("\nGoroutine stopped")
				return
			}
		}
	}()

	time.Sleep(3 * time.Second)
	stop <- true

	// Даем время для вывода сообщения о завершении горутины
	time.Sleep(1 * time.Millisecond)
}

// 2. Использование общих переменных для сигнализации завершения
func methodTwo() {
	// По истечении времени меняем значение в переменной ok на false,
	// тем самым заходя в блок if !ok { ... } в котором мы завершаем
	// выполнение горутину вызовом return

	var wg sync.WaitGroup
	ok := true

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			if !ok {
				fmt.Println("\nGoroutine stopped")
				return
			}

			fmt.Println("Goroutine is running...")
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(3 * time.Second)
	ok = false
	wg.Wait()

	// Даем время для вывода сообщения о завершении горутины
	time.Sleep(1 * time.Millisecond)
}

// 3. Использование закрытия канала в качестве сигнала завершения
func methodThree() {
	// По истечении времени закрываем канал stop, тем самым
	// вызывая секцию stop в операторе select { ... }

	stop := make(chan struct{})

	go func() {
		for {
			select {
			default:
				fmt.Println("Goroutine is running...")
				time.Sleep(time.Second)
			case <-stop:
				fmt.Println("\nGoroutine stopped")
				return
			}
		}
	}()

	time.Sleep(3 * time.Second)
	close(stop)

	// Даем время для вывода сообщения о завершении горутины
	time.Sleep(1 * time.Millisecond)
}

// 4. Использование контекста для управления жизненным циклом горутины
func methodFour() {
	// По истечении времени вызываем метод cancel
	// и переходим в блок <-ctx.Done() оператора switch { ... }
	// тем самым завершая выполнение горутины с помощью return
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("\nGoroutine stopped")
				return
			default:
				fmt.Println("Goroutine is running...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()

	// Даем время для вывода сообщения о завершении горутины
	time.Sleep(1 * time.Millisecond)
}

func Run() {
	methodOne()
	methodTwo()
	methodThree()
	methodFour()
}
