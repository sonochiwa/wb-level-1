package exercise_25

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func sleep(d time.Duration) {
	// Канал будет закрыт после временного интервала d
	<-time.After(d)
}

func Run() {
	fmt.Printf("Легли спать...\n")
	sleep(5 * time.Second)
	fmt.Printf("...проснулись.")
}
