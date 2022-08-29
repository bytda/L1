package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

/*Задание:
Реализовать собственную функцию sleep.
*/

func main() {
	var duration int
	fmt.Println("Введите длительность в секундах: ")
	fmt.Fscan(os.Stdin, &duration) //получаем длительность от пользователя

	SleepTimer(time.Second * time.Duration(duration)) //запускаем в секундах
	SleepCTX(time.Second * time.Duration(duration))

}
func SleepTimer(duration time.Duration) { //способ с помощью таймера
	timer := time.After(duration) //задаем таймер продолжительность получаем из параметров функции
	fmt.Println("Начинаем засыпать....Timer")
	<-timer //ждем сигнала завершения таймера
	fmt.Println("Проснулись")
}
func SleepCTX(duration time.Duration) { //способ с помощью контекста
	ctx, _ := context.WithTimeout(context.Background(), duration) //задаем контекст с таймером
	fmt.Println("Начинаем засыпать....CTX")
	<-ctx.Done() //ждем сигнал завершения
	fmt.Println("Проснулись")
}
