package exercise_7

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func writeToMap(likes map[string]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	likes["counter"] += 1
	mu.Unlock()
}

func Run() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	likes := map[string]int{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go writeToMap(likes, &wg, &mu)
	}

	wg.Wait()

	fmt.Println(likes["counter"])
}
