package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Задание:
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
func main() {
	var n int //количество воркеров

	fmt.Printf("Введите необходимое количество воркеров ") //выводим в stdout обращение к пользователю
	_, err := fmt.Scan(&n)                                 //получаем данные введенные от пользователя в консоль
	if err != nil {                                        //в случае ошибки, выводим ее ошибку в консоль
		fmt.Println(err)
		return
	}

	mainChannel := make(chan int, 100) //объявляем канал (главный поток)
	defer close(mainChannel)           //после выполнения функции main() закрываем канал

	ctx, cancel := context.WithCancel(context.Background()) //объявляем контекст с функцией отмены

	finishChannel := make(chan bool, n) // объявляем канал для контроля завершения всех воркеров
	defer close(finishChannel)

	for i := 0; i < n; i++ { //запускаем n количество воркеров
		go Worker(ctx, finishChannel, mainChannel, i)
	}
	go ExitSignal(cancel) //запускаем функцию перехвата сигнала о том что надо закрыть программу
loop: //даем лейбл циклу чтобы потом завершить его
	for {
		select {

		case mainChannel <- rand.Intn(10): //отправляем в главный канал случайные числа от 0 до 9
			time.Sleep(time.Second) // немного замедляем выполнение
		case <-ctx.Done():
			break loop // выходим из цикла при получении сигнала от контекста
		}
	}
	for i := 0; i < n; i++ {
		<-finishChannel //ожидаем выполнения всех воркеров
	}
	os.Exit(0) //завершаем программу с кодом 0

}
func Worker(ctx context.Context, finishChan chan<- bool, mainChan <-chan int, id int) {
	//используем лейбл чтобы потом выйти из цикла
loop2:
	for {
		select {
		case info := <-mainChan:
			fmt.Println(id, "-ий Воркер получил:", info) //если получаем данные из канала
			// - выводим в консоль
		case <-ctx.Done():
			break loop2 // если получаем сигнал выхода, выходим из цикла
		}
		fmt.Println("Воркер номер ", id, " закончил работу") //отправляем в канал данные что работа воркера окончена
		finishChan <- true
	}

}
func ExitSignal(cancel context.CancelFunc) {
	a := make(chan os.Signal, 1)                      //объявим канал для отправки сигнала о завершении программы
	signal.Notify(a, syscall.SIGINT, syscall.SIGTERM) //перехватываем сигнал
	fmt.Println("Получаем сигнал о завершении", <-a)

	cancel() // вызываем функцию завершения в контексте
}
