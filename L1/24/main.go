package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) GetX() float64 { // геттер
	return p.x
}

func (p *Point) GetY() float64 { // сеттер
	return p.y
}

func NewPoint(x, y float64) *Point { // конструктор с x,y для x,y
	return &Point{x: x, y: y}
}

func Distance(p1, p2 *Point) float64 { // вычисление расстояния между точками (двумя)
	return math.Sqrt(math.Pow(p2.GetX()-p1.GetX(), 2) + math.Pow(p2.GetY()-p1.GetY(), 2))
}

func main() {
	p1 := NewPoint(1.0, 2.0)
	p2 := NewPoint(4.0, 6.0)
	distance := Distance(p1, p2)
	fmt.Println(distance)
}
