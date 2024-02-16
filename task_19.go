package main

/*
Задача:
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

	Логика:
	1) Просто берем символы с конца и ставим в начало.
*/

func reverseString(s string) string {
	runes := []rune(s)
	var reversed []rune
	for i := len(runes) - 1; i >= 0; i-- {
		reversed = append(reversed, runes[i])
	}
	return string(reversed)
}
