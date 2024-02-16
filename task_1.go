package main

import "fmt"

/*

	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

*/

// Human - основная структура
type Human struct {
	//Встраивание структуры Action
	Action
}

type Action struct {
}

// Метод для примера встраивания
func (act *Action) HelloWorld() {
	fmt.Println("Hello world")
}

func task_1() {

	var human Human = Human{}
	human.HelloWorld()
}
