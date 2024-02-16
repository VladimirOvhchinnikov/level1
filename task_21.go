package main

/*
Задача:
Реализовать паттерн «адаптер» на любом примере.

У нас есть интерфейс OldNotifier, который требует иметь для его реализации метод Send.
MyEmailNotifier реализует его. Будем считать их как комплекс старой реализации
Так же у нас есть новый интерфейс Notifier.
И реализация этого интерфейса с оберткой



	Логика:
	1)Старый интерфейс OldNotifier помещается в структуру EmailNotifierAdapter => любая реализация,
	реализующая данный интерфейс, может быть в этой структуре
	2)Notify структуры EmailNotifierAdapter делегиурет вызов методу SendEmail старого интерфейса EmailNotifier
*/

type OldNotifier interface {
	Send(message string) string
}

type MyEmailNotifier struct{}

func (n *MyEmailNotifier) Send(message string) string {
	return "send" + message
}

type Notifier interface {
	Notify(message string) string
}

type EmailNotifierAdapter struct {
	OldNotifier OldNotifier
}

func (a *EmailNotifierAdapter) Notify(message string) string {
	return a.OldNotifier.Send(message)
}
