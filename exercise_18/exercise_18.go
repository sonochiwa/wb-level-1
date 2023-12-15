package exercise_18

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.

type Counter struct {
	quantity int
	mu       sync.Mutex
}

// NewCounter - конструктор счетчика
func NewCounter() *Counter {
	return &Counter{
		quantity: 0,
	}
}

// Increment - метод инкрементирующий счетчик
func (c *Counter) Increment() {
	c.mu.Lock()         // Блокируем данные, для безопасной записи
	defer c.mu.Unlock() // Разблокируем
	c.quantity += 1
}

func Run() {
	var wg sync.WaitGroup
	counter := NewCounter()

	// Запускаем 500 горутин, которые увеличивают счетчик
	for i := 0; i < 500; i++ {
		wg.Add(1)

		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.quantity) // quantity совпадает с количеством запущенных горутин
}
