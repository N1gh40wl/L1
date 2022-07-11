package task24

import (
	"fmt"
	"math"
)

type Point struct {
	x float32
	y float32
}

func NewPoint(x float32, y float32) Point {
	p := Point{
		x: x,
		y: y,
	}
	return p
}

func length(a, b Point) float32 {
	return float32(math.Sqrt(math.Pow((float64(b.x)-float64(a.x)), 2) + math.Pow((float64(b.y)-float64(a.y)), 2)))
}

func RunTask() {
	var x, y float32
	fmt.Println("Введите координаты 1 точки:")
	fmt.Scanf("%f %f", &x, &y)
	a := NewPoint(x, y)
	fmt.Println("Введите координаты 2 точки:")
	fmt.Scanf("%f %f", &x, &y)
	b := NewPoint(x, y)
	fmt.Printf("Расстояние между точками: %f\n", length(a, b))

}
