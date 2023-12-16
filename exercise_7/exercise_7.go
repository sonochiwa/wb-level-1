package exercise_7

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func writeToMap(likes map[string]int, wg *sync.WaitGroup, mu *sync.RWMutex) {
	// Следуя хорошим практикам вызываем wg.Done() используя отложеный вызов
	defer wg.Done()

	mu.Lock() // Блокируем
	likes["counter"] += 1
	mu.Unlock() // Разблокируем
}

func Run() {
	var wg sync.WaitGroup
	var mu sync.RWMutex

	likes := map[string]int{} // Карта, в которой будем считать лайки

	// Запускаем 100 горутин изменяющих нашу мапу
	for i := 0; i < 100; i++ {
		wg.Add(1) // Добавляем 1
		go writeToMap(likes, &wg, &mu)
	}

	wg.Wait() // Ждем обнуления счетчика

	fmt.Println(likes["counter"]) // 100 - вывод соответствует ожиданиям
}
