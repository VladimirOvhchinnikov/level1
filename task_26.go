package main

import "strings"

/*
Задача:
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false

Для решения будет взята логика из 12 задачи.

	Логика:
	1)на вход приходит строка
	2) создаем мапу, где значение - булевая переменная
	3) Проходимся по всем символам, которые приведены к малому знаку
	4) на каждой иттерации проверяем существует ли символ. Если да, то возвращаем false, если нет, то сохраняем в мапу

*/

func areAllCharactersUnique(str string) bool {
	charSet := make(map[rune]bool)
	for _, char := range strings.ToLower(str) {
		if _, exists := charSet[char]; exists {
			return false
		}
		charSet[char] = true
	}
	return true
}
