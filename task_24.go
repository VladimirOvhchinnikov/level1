package main

import "math"

/*
Задача:
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.


	Логика:
	1)
*/

// Структура Point, которая инкапсулирует
type Point struct {
	x float64
	y float64
}

// Простой конструктор
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// функция расчета дистанции
func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}
