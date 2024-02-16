package main

import (
	"fmt"
	"math/big"
	"reflect"
)

/*
Задача:
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

	Я решил, что проще всего будет добавить логику конвертации числа из любого вида и возвращать big.Int
	convertToBigInt принимает любой тип данных. Тут можно добавить логику для проверки любых пременных


	Логика:
	1)Числа конвертируются в convertToBigInt в big.Int
	2)Дальше используется одна из функций к примеру addNumbers
*/

func convertToBigInt(num interface{}) (*big.Int, error) {
	if v, ok := num.(int64); ok {
		return big.NewInt(v), nil
	} else {
		return nil, fmt.Errorf("unsupported type: %s", reflect.TypeOf(num))
	}
}
func addNumbers(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func subtractNumbers(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func multiplyNumbers(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func divideNumbers(a, b *big.Int) (*big.Int, error) {
	if b.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	return new(big.Int).Div(a, b), nil
}
