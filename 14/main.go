package main

import (
	"fmt"
	"reflect"
)

/*
Задание:
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	var val interface{} //подопытная переменная
	var str string      //переменная для вывода

	val = "hp" //присваиваем переменной тип
	str = Switch(val)
	fmt.Println("Результат Switch:", str)
	str = Reflect(val)
	fmt.Println("Результат Reflect:", str)
	str = Fmt(val)
	fmt.Println("Результат Fmt:", str)
}

func Switch(obj interface{}) string { //метод с помощью Switch
	var str string      //создаем переменную для вывода
	switch obj.(type) { //используем switch по типу элементу
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan bool:
		return "chan bool"
	case chan string:
		return "chan string"

	}
	return str
}
func Reflect(obj interface{}) string { //с помощью Reflect
	var str string                     //создаем переменную для вывода
	str = reflect.TypeOf(obj).String() //используем reflect чтобы получить тип объекта
	return str
}

func Fmt(obj interface{}) string { // с помощью fmt
	var str string               //создаем переменную для вывода
	str = fmt.Sprintf("%T", obj) //используем форматирование пакета fmt
	return str
}
