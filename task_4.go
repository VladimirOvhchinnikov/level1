package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
 Задача:
 Реализовать постоянную запись данных в канал (главный поток).
 Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
 Необходима возможность выбора количества воркеров при старте.

 Причины такой реализации:
 - объектное восприятие для человека часто проще,
 - возможность повторного использования - позволяет сократить расходы на инициализацию рутин и их закрытие

 Логика работы:
 - Есть канал. В него отправляются задачи
 - Есть работники. Они постоянно считывают и берут в работу задачи

 Graceful Shutdown - постарался приблизиться к данному шаблону заершения.
*/

// Структура работа.
// Может хранить любую информацию, которая характерезует суть абстрактной деятельности
// В моем случае - просто текст
type Job struct {
	data string
}

// Структура работник.
// Может хранить любую информацию, которая характерезует суть абстрактного работника, но должна
// иметь номер индентифиактор, канал для считывания работ, структуру ожидания.
// В моем случае - только стандартные переменные
type Worker struct {
	id   int
	jobs <-chan Job
	wg   *sync.WaitGroup
}

// Функция-конструктор работника
func NewWorker(id int, jobs <-chan Job, wg *sync.WaitGroup) Worker {
	return Worker{
		id:   id,
		jobs: jobs,
		wg:   wg,
	}
}

// Метод структуры Worker
// Запускает рутину для выполнения задач из канала jobs
func (w Worker) Start(stopChan <-chan struct{}) {

	//Зпуск рутины
	go func() {
		//Бесконечный цикл, пока не будет выполнено одно из условий закрытия
		for {
			//Два кейса
			// - когда закрытие происходит из за закрытия канала
			// - когда получен сигнал от канала stopChan
			select {
			case job, ok := <-w.jobs:
				if !ok {
					return // Канал закрыт, завершаем работу воркера
				}
				// Обработка задачи
				fmt.Println(job.data)
				w.wg.Done()
			case <-stopChan:
				fmt.Println("ctrl + c")
				return // Получен сигнал остановки, завершаем работу воркера
			}
		}
	}()
}

func Test_4() {

	//Создаем константы для воркров и работ
	const numbersJobs = 500000
	const numbersWorkers = 3

	//Инициализируем канал
	jobs := make(chan Job, numbersJobs)
	var wg sync.WaitGroup

	//Инициализируем канал остановки
	stopChan := make(chan struct{})
	//Инициализируем канал для считывания сигнала
	sigChan := make(chan os.Signal, 1)
	//прописываем отслеживание  Ctrl+C
	signal.Notify(sigChan, syscall.SIGINT)

	//Запускаем воркера
	for i := 0; i < numbersWorkers; i++ {
		worker := NewWorker(i, jobs, &wg)
		worker.Start(stopChan)
	}

	go func() {
		<-sigChan
		close(stopChan) // Посылаем сигнал остановки воркерам
		close(jobs)     // Закрываем канал заданий
	}()

	//Запускаем в каналы задания
	for i := 1; i <= numbersJobs; i++ {
		wg.Add(1)
		jobs <- Job{data: fmt.Sprint(i)}
	}

	// закрываем канал
	close(jobs)
	// дождиаемся окончания работ
	wg.Wait()
}
