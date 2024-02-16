package main

import "reflect"

/*
 Задача:
 Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

/*
Логика работы.
  Можно было бы использовать v.(type), но лучшим решением будет использовать пакет reflect
  В силу более мощного инструментария

  1)Получаем значение
  2)ЧЕрез свитч сравниваем
*/

func test_14(v interface{}) string {
	// Получение значения reflect.Type переменной
	t := reflect.TypeOf(v)

	// Проверяем не пусто ли
	if t == nil {
		return "nil"
	}

	// Возвращение строки с именем типа
	switch t.Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "channel"
	default:
		return "unknown"
	}
}
