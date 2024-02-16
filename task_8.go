package main

import (
	"fmt"
	"strconv"
)

/*
 Задача:
 Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func Test_8(number int64, index int64, numberPosition int) int64 {
	var chars []rune

	// Цикл где мы число переводим в строку и потом записываем в срез рун
	for _, char := range fmt.Sprintf("%064b", number) {
		chars = append(chars, char)
	}

	// Проверяем, что пришел запрос на изменение в диапозоне 64
	if numberPosition < 0 || numberPosition >= 64 {
		fmt.Println("Ошибка: numberPosition вне диапазона")
		return number
	}

	// Установка нужного бита
	if index == 1 {
		/// Установить бит в 1
		chars[63-numberPosition] = '1'
	} else {
		//Установить бит в 0
		chars[63-numberPosition] = '0' //
	}

	// Преобразование обратно в int64
	result, err := strconv.ParseInt(string(chars), 2, 64)
	if err != nil {
		fmt.Println("Ошибка при преобразовании:", err)
		return 0
	}
	return result
}

func SetBit(number int64, position int, value int64) int64 {
	if value == 0 {
		// Установить бит в 0
		return number &^ (1 << position)
	}
	// Установить бит в 1
	return number | (1 << position)
}
