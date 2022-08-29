package main

import (
	"math/rand"
	"sync"
)

/*
Задание:
Реализовать конкурентную запись данных в map.
*/
type Data struct { //объявляем вспомогательную структуру
	id   int
	name int
}

func main() {
	//Channel()
	Mutex()
}

func Channel() {
	n := 5                         // количество писателей
	data := make(map[int]int)      //создаем мапу для записи
	chanData := make(chan Data, n) //создаем канал для передачи данных

	for i := 0; i < n; i++ { //запускаем n горутин

		go func(ch chan Data) {
			for {
				ch <- Data{id: i, name: i} //отправляем данные в канал
			}
		}(chanData)

	}
	for { //бесконечный цикл чтения из канала
		tmp := <-chanData //получаем данных
		data[tmp.id] = tmp.name
	}
}

func Mutex() {
	n := 5 //задаем количество писателей

	data := make(map[int]int) //задаем мапу

	var mtx sync.Mutex //создаем Mutex

	for i := 0; i < n; i++ {
		go func(data map[int]int, mtx *sync.Mutex) {
			for {
				id := rand.Int()
				name := rand.Int()
				mtx.Lock()      //блокируем Mutex
				data[id] = name //записываем
				mtx.Unlock()    //открываем Mutex
			}
		}(data, &mtx)

	}

}

func SyncMap() {
	n := 5                   //задаем количество писателей
	var data sync.Map        //создаем мапу с
	for i := 0; i < n; i++ { //начинаем цикл
		go func(data *sync.Map) {
			for {
				data.Store(rand.Int(), rand.Int()) //записываем значения в мапу
			}
		}(&data)
	}
}
