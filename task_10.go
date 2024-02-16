package main

import (
	"fmt"
	"math"
)

/*
 Задача:
 Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
 Объединить данные значения в группы с шагом в 10 градусов.
 Последовательность в подмножноствах не важна.
*/

/*
Логика работы.
 1) groupedTemperatures получает на вход срез чисел и длину шага
 2) прозодимся по каждому значению
 3) отправляем getRangeValue значение для определения диапозона
 4) полученные данные добвляем в соответствующей срез в карте
*/

func test_10() {
	//Создаем  срез чесел
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//отправялем в фнукцию, которая прнимает срез и длину интервала
	groupedTemperatures := groupTemperatures(temperatures, 10)

	//Выводим результат работы
	for rangeVal, temps := range groupedTemperatures {
		fmt.Printf("%v: %v\n", rangeVal, temps)
	}

}

// Функция groupTemperatures принимает срез температур (temps) и числовой шаг (step).
func groupTemperatures(temps []float64, step float64) map[string][]float64 {
	// Создаем карту для хранения группированных температур.
	groups := make(map[string][]float64)

	// Проходим по каждой температуре в срезе.
	for _, temp := range temps {
		// Получаем строковое представление диапазона для текущей температуры
		rangeVal := getRangeValue(temp, step)

		// Добавляем текущую температуру в соответствующий срез в карте
		groups[rangeVal] = append(groups[rangeVal], temp)
	}
	return groups
}

// Функция getRangeValue определяет диапазон для заданной температуры.
func getRangeValue(temp, step float64) string {
	// Вычисляем нижнюю границу диапазона и делим температуру на шаг
	lower := math.Floor(temp/step) * step

	// Вычисляем верхнюю границу диапазона как нижнюю границу плюс шаг
	upper := lower + step

	return fmt.Sprintf("[%v, %v)", lower, upper)
}
