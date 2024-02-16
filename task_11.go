package main

import (
	"fmt"
	"sort"
)

/*
 Задача:
 Реализовать пересечение двух неупорядоченных множеств.
*/

/*
Логика работы.
 1) Я выбрал решение в лоб, потому что задание не раскрыто в полной мере
 2) Берем множества и сортируем их
 3) Двумя циклами ищем одинаковые значения и добавляем их в intersection
 4) Выводим результат
*/

func test_11() {
	//Срезы
	set1 := []int{5, 3, 4, 1, 2}
	set2 := []int{8, 7, 6, 5, 4}

	//Вызов функции нахождения пересечений
	intersection := intersectSortedWithLoops(set1, set2)
	fmt.Println("Пересечение множеств:", intersection)
}

// Функция для нахождения пересечения двух отсортированных множеств с использованием двух циклов
func intersectSortedWithLoops(set1, set2 []int) []int {

	//Сортируем
	sort.Ints(set1)
	sort.Ints(set2)

	var intersection []int

	//Проходимся циклами для сравнения
	for _, item1 := range set1 {
		for _, item2 := range set2 {
			if item1 == item2 {
				// Проверяем, не является ли текущий элемент дубликатом последнего добавленного элемента в пересечение
				if len(intersection) == 0 || intersection[len(intersection)-1] != item1 {
					intersection = append(intersection, item1)
				}
				break
			}
		}
	}

	return intersection
}
