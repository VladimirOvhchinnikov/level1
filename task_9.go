package main

import "fmt"

/*
 Задача:
 Разработать конвейер чисел. Даны два канала:
 в первый пишутся числа (x) из массива, во второй — результат операции x*2,
 после чего данные из второго канала должны выводиться в stdout.
*/

func Test_9() {

	numbers := []int{1, 2, 3, 4, 5}

	// Создаем каналы
	input := make(chan int)
	output := make(chan int)

	// Горутина для чтения чисел и записи их в первый канал
	go func() {
		for _, x := range numbers {
			input <- x
		}
		close(input)
	}()

	// Горутина для чтения из первого канала, удвоения чисел и записи их во второй канал
	go func() {
		for x := range input {
			output <- x * 2
		}
		close(output)
	}()

	// Чтение из второго канала и вывод в stdout
	for result := range output {
		fmt.Println(result)
	}
}
