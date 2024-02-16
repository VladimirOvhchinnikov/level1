package main

import "sort"

/*
Задача:
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

Встроенные методы языка go связанные с сортировкой  быстрой - sort.Ints
Но думаю, что задача стояла написать сортировку самому.

1) Мы можем использовать sort.Ints или ему подобные решения
2) Мы можем написать реализацию для sort.Interface и использовать взаимозсвяннаые наработки
3) Мы можем применить логику полность на базе своих решений оставив только алгоритм

*/

// Первая версия просто использует готовое решение
func version1(slice []int) {
	sort.Ints(slice)
}

// Приведем в пример наш выдуманный тип данных
type MyData []int

// Реализация методов интерфейса sort.Interface
// Len() возвращает  длину в нашем наборе данных
func (d MyData) Len() int { return len(d) }

// Less логика проверки кто за кем стоит
func (d MyData) Less(i, j int) bool { return d[i] < d[j] }

// Swap алгоритм разворота
func (d MyData) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

//после чего мы можем вызывать  sort.Sort(MyData)

// функция для разделения массива
func partition(arr []int, low, high int) int {
	// Выбираем опорный элемент как последний элемент в массиве
	pivot := arr[high]

	// Индекс меньшего элемента
	i := low

	// Проходимся по всем элементам массива
	for j := low; j < high; j++ {
		// Если текущий элемент меньше или равен опорному
		if arr[j] < pivot {
			// Меняем местами arr[i] и arr[j]
			arr[i], arr[j] = arr[j], arr[i]
			// Перемещаем указатель на один элемент вперед
			i++
		}
	}

	// Помещаем опорный элемент между элементами меньшими и большими от него
	arr[i], arr[high] = arr[high], arr[i]
	return i // Возвращаем индекс опорного элемента
}

// функция быстрой сортировки
func quickSort(arr []int, low, high int) {
	// Базовый случай: если есть хотя бы два элемента для сортировки
	if low < high {
		// Разделяем массив и получаем индекс опорного элемента
		pi := partition(arr, low, high)

		// Рекурсивно сортируем элементы до опорного элемента
		quickSort(arr, low, pi-1)
		// Рекурсивно сортируем элементы после опорного элемента
		quickSort(arr, pi+1, high)
	}
}
