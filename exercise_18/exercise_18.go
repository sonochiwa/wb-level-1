package exercise_18

// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.

type Counter struct {
	quantity int
}

// NewCounter - конструктор счетчика
func NewCounter() *Counter {
	return &Counter{
		quantity: 0,
	}
}

// Increment - метод инкрементирующий счетчик
func (c *Counter) Increment() {
	c.quantity += 1
}

func Run() {

}
