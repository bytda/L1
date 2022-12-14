package main

import (
	"fmt"
	"sync"
)

/*
Задание:
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
func main() {

	Channels() //вызываем функции
	WaitGroup()
}

func Channels() {
	fmt.Println("С использованием каналов:") // выводим в консоль заголовок первого решения задачи

	var arr = [...]int{2, 4, 6, 8, 10} // создаем массив с заданными числами

	structCh := make(chan struct{}, len(arr)) // создаем буферизированный канал выполняющий роль синхронизатора,
	//чтобы все горутины выполнились до завершения main()

	for _, val := range arr { // перебираем элементы массива и передаем значения в функцию вычисления квадрата числа

		// запускаем функцию в отдельной горутине
		go SqrChannels(structCh, val)
	}

	for i := 0; i < len(arr); i++ { // ждем завершения работы горутин путем считывания из канала кол-ва результатов,
		// равное количеству запущенных горутин

		<-structCh
	}
}

func SqrChannels(ch chan<- struct{}, val int) { //объявляем канал ch только для отправки
	fmt.Println(val, "*", val, "=", val*val) // выводит в stdout оформленный результат вычисления квадрата val
	ch <- struct{}{}                         // отправляем в канал данные, отправляем пустую структуру так как она самая легкая

}
func WaitGroup() {
	fmt.Println("С использованием пакета sync:")

	var arr = [...]int{2, 4, 6, 8, 10} // создаем массив с заданными числами
	var wg sync.WaitGroup              // создаем экземпляр sync.WaitGroup для контроля окончания выполнения группы горутин
	// чтобы все горутины выполнились до завершения main()
	wg.Add(len(arr)) // передаем в wg число горутин в группе, завершения которых мы будем дожидаться,
	// то есть сколько вызовов Done() мы должны отследить

	for _, val := range arr { // перебираем элементы массива и передаем значения в функцию вычисления квадрата числа

		go SqrWG(&wg, val) // запускаем функцию в отдельной горутине

	}
	wg.Wait() // ждем сигналы завершения Done() от всех горутин в группе

}

func SqrWG(wg *sync.WaitGroup, val int) {
	fmt.Println(val, "*", val, "=", val*val) // выводит в stdout оформленный результат вычисления квадрата val

	defer wg.Done() // по окончанию выполнения функции отправляет в wg сигнал,
	// что все необходимые операции были выполнены
}
