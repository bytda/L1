package main

import "fmt"

/*
Задача: Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).
*/

type Human struct { // объявляем родительскую структуру
	name string
	age  int
}

func (p Human) Hello() string { //объявляем метод родительской структуры
	return fmt.Sprintf("Hi, I am %s", p.name)
}

type Action struct { //встраиваем структуру Human в структуру Action
	Human
}

func main() {

	action := Action{Human{"John", 25}} //объявляем объект наследника Action
	fmt.Println(action.Hello())         //выводим в консоль строку

}
