package main

import (
	"fmt"
	"math"
)

/*
Задание:
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x, y float64
}

func main() {
	point1 := NewPoint(0, 0) // создаем две точки
	point2 := NewPoint(4, -7)

	fmt.Println("Расстояние между двумя точками:", PointDistation(point1, point2)) // вычисляем расстояние между ними и выводим его в stdout
}

func NewPoint(x, y float64) *Point { // NewPoint конструктор структуры Point
	return &Point{x: x, y: y} // задаем значения полей и возвращаем указатель на созданную структуру
}

func PointDistation(p1, p2 *Point) float64 { //вычисляет расстояние между point1 и point2
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)) // вычисляем расстояние по формуле
}
