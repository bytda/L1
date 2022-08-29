package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Задание:
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	//CxtTimeOut() //способ с помощью контекста с таймером
	//CtxCancel() // с помощью контекста с отменой
	//Timer() //  с помощью таймера
	//Channel() // с помощью канала
	//Mutex() //с помощью Mutex
	//WaitGroup() //с помощью WaitGroup
}

func CxtTimeOut() {
	finishChannel := make(chan struct{})                               //канал для отслеживания завершения
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5) //контекст который завершится через 5 сек

	go func() { //запускаем горутину
	loop:
		for { //запускаем бесконечный цикл
			select {
			case <-time.Tick(time.Second): //каждую секунду
				fmt.Println("tik..") // пишем в консоль
			case <-ctx.Done(): //как получили сигнал от ctx
				break loop // прерываем цикл
			}
		}
		fmt.Println("Ctx убил горутину") //пишем в консоль
		finishChannel <- struct{}{}      //отправляем сигнал в канал отслеживающий завершение
	}()
	<-finishChannel //ждем сигнал о завершении
}
func CtxCancel() {
	finishChannel := make(chan struct{})                    //канал для отслеживания завершения
	ctx, cancel := context.WithCancel(context.Background()) //контекст с функцией отмены

	go func() { //запускаем горутину
	loop2: //объявляем лейбл циклу чтобы завершить его
		for {
			select {
			case <-time.Tick(time.Second):
				fmt.Println("..tic") //каждую секунду отправляем сообщение в stdout
			case <-ctx.Done():
				break loop2 //  получив сигнал от контекста выходим из цикла
			}
		}
		fmt.Println("CtxCancel убил горутину")
		finishChannel <- struct{}{} //отправляем сигнал что горутина завершилась
	}()
	time.Sleep(time.Second * 6) //задерживаем жизнь горутины на 6 секунды
	cancel()                    //отправляем сигнал завершения контексту
	<-finishChannel             //ждем сигнал завершения от горутины
}
func Timer() {
	secondsBeforeKill := 5               //определяем время для горутины
	finishChannel := make(chan struct{}) //канал для отслеживания заверешения горутины
	go func() {
		timer := time.After(time.Duration(secondsBeforeKill) * time.Second) //задаем таймер завершения горутины
	loop3: //лейбл для цикла
		for {
			select {
			case <-time.Tick(time.Second): //каждую секунду
				fmt.Println("..tic") //отправляем сообщение
			case <-timer: //получив сигнал от таймера
				break loop3 //прерываем цикл
			}
		}
		fmt.Println("Timer убил горутину")
		finishChannel <- struct{}{} //отправляем сигнал о завершении горутины
	}()
	<-finishChannel //ждем закрыти горутины
}
func Channel() {

	killerChan := make(chan struct{}) //канал для отправки сигнала завершения
	finishChan := make(chan struct{}) //канал для отслеживания завершения горутины

	go func() { //запускаем горутину
	loop4: //лейбл для цикла
		for {
			select {
			case <-time.Tick(time.Second): //каждую секунду пишем в stdout
				fmt.Println("tic..")
			case <-killerChan: //как только получили сигнал о завершении
				break loop4 //выходим из цикла
			}
		}
		fmt.Println("Channel убил горутину") //пишем в консоль
		finishChan <- struct{}{}             //отправляем в канал сигнал что горутина закончила
	}()
	time.Sleep(time.Second * 4) //ждем 4 секунды перед отправкой сигнала о завершении горутины
	killerChan <- struct{}{}    //отправляем сигнал о завершении(отправляем пустую структуру так как она самая легкая)
	<-finishChan                //ожидаем завершение горутины
}
func Mutex() {
	var mtx sync.Mutex //объявляем экземпляр Mutex
	go func() {        //запускаем горутину
		for {
			select {
			case <-time.Tick(time.Second): //каждую секунду
				fmt.Println("tic..") //пишем в консоль

			}
			mtx.Lock()   //блокируем
			mtx.Unlock() //и разблокируем Mutex
		}

	}()
	time.Sleep(time.Second * 3)        //ждем 3 секунды до завершения горутины
	mtx.Lock()                         //блоикируем Mutex, так как мы заблокировали его здесь, в цикле Mutex не сможет закрыться и завершит цикл
	fmt.Println("Mutex убил горутину") //пишем в консоль
}

func WaitGroup() {
	var wg sync.WaitGroup //объявляем экземпляр WaitGroup
	go func() {           //запускаем горутину
		for { //бесконечный цикл
			select {
			case <-time.Tick(time.Second): //каждую секунду
				fmt.Println("tic..") //отправляем в консоль тик

			}
			wg.Wait() //WaitGroup ожидает завершения, но так как в группе 0, горутин он не блокирует функцию
		}
	}()
	time.Sleep(time.Second * 4) //ждем 4 секунды перед завершением горутины
	wg.Add(1)                   //добавляем в группу одну горутину, из-за этого блокируется цикл,
	// так как он ждет завершения группы но Done() нигде мы не указывали
	fmt.Println("WaitGroup убил горутину") //пишем в консоль
}
