package main

import (
	"fmt"
	"strings"
)

/*
Задание:
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	str := "snow dog sun"                  //объявляем начальную строку
	fmt.Println("Начальная строка: ", str) //выводим ее
	strSlice := strings.Split(str, " ")    //делим строку на слова и создаем из них слайс
	Reverse(strSlice)                      //вызываем функцию переворота
	fmt.Println("Результат: ", strSlice)   //выводим результат
}
func Reverse(data []string) { //функция переварачивания
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 { //перебираем слайс с обеих сторон
		data[i], data[j] = data[j], data[i] //меняем местами элементы
	}

}
