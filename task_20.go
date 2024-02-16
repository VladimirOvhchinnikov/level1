package main

import "strings"

/*
Задача:
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».


	Логика:
	1) Разделем предложение на срез строк.
	2) Записываем значения в обратном порядке
	3) Возвращаем новую строку
*/

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	var reversedWords []string

	for i := len(words) - 1; i >= 0; i-- {
		reversedWords = append(reversedWords, words[i])
	}

	return strings.Join(reversedWords, " ")
}
