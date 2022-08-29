package main

import (
	"context"
	"fmt"
	"time"
)

/*
Задание:
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
func main() {
	var n int //количество секунд
	fmt.Print("Введите необходимое количество секунд: ")
	_, err := fmt.Scan(&n) //получаем от пользователя нужное количество секунд
	if err != nil {
		fmt.Println(err) // обрабатываем ошибку
		return
	}
	finishChannel := make(chan int) // создаем канал для сигнала завершения работы
	go Ctx(n, finishChannel)        //решение задачи с помощью контекста
	go Timer(n, finishChannel)      //решение задачи с помощью таймера

loop:

	for {
		select {
		case <-time.Tick(time.Second):
			fmt.Println("..sec")
		case <-finishChannel:
			break loop
		}
	}

}

func Ctx(n int, finishChannel chan int) {
	//создаем контекст с таймером самозавершения
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(n))

	intChan := make(chan int)  //инициализируем канал
	go WriterCtx(ctx, intChan) //запускаем писателя в отдельной горутине

loop2: //цикл чтения из канала
	for {
		select {
		case a, ok := <-intChan: //читаем канал и проверяем открыт ли он
			if ok {
				fmt.Println("Ctx получил:", a) //если открыт , пишем в консоль
				break
			}
			fmt.Println("Чтение Ctx прекращено") //если закрыт, выходим из цикла чтения
			break loop2
		}
	}
	close(finishChannel) //посылаем сигнал о завершении работы функции закрытием канала
}
func WriterCtx(ctx context.Context, intChan chan int) { //функция писателя в канал
	var x, y = 1, 0
loop2:
	for {
		y, x = x, y+x
		select {
		case intChan <- y:
			time.Sleep(time.Millisecond * 100) //замедляем, чтобы не засорять консоль
		case <-ctx.Done(): //если получили сигнал закрытого контекста
			fmt.Println("Писатель завершился")
			close(intChan) //закрываем канал
			break loop2    //прерываем цикл

		}
	}
}
func Timer(n int, finishChan chan int) {
	intChan := make(chan int)  //инициализируем канал
	go WriterTimer(intChan, n) //запускаем писателя в отдельной горутине

loop2:
	for {
		select {
		case a, ok := <-intChan:
			if ok { //если канал открыт - выводим данные в консоль
				fmt.Println("Таймер получил: ", a)
				break
			}
			fmt.Println("Таймер завершил работу") //если канал закрыт, завершаем работу
			break loop2                           //выходим из цикла чтения

		}
	}
	close(finishChan)
}

func WriterTimer(intChan chan int, n int) {
	var x, y = 1, 0
	timer := time.After(time.Second * time.Duration(n)) //задаем таймер
loop3:
	for {
		y, x = x, x+y //генерируем числа фибоначи
		select {
		case intChan <- y: //отправляем данные в канал
			time.Sleep(time.Millisecond * 100) //с задержкой
		case <-timer: //если пришел сигнал от таймера что время истекло
			fmt.Println("Таймер писатель закончил работу") //пишем в консоль
			close(intChan)                                 //закрываем канал
			break loop3                                    //выходим из цикла писателя
		}
	}
}
