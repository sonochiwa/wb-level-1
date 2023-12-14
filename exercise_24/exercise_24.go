package exercise_24

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.

type Point struct {
	X, Y float64
}

// Конструктор для создания точки
func NewPoint(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// Метод для вычисления расстояния между двумя точками
func (p1 Point) DistanceTo(p2 Point) float64 {
	distanceX := p1.X - p2.X
	distanceY := p1.Y - p2.Y
	return math.Sqrt(distanceX*distanceX + distanceY*distanceY)
}

func Run() {
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	distance := point1.DistanceTo(point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
