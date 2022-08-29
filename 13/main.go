package main

/*Задание:
Поменять местами два числа без создания временной переменной.
*/

func main() {
	var a, b = 1, 2
	println("a:", a, "b:", b)

	a, b = b, a //способ с помощью возможностей языка
	println("a:", a, "b:", b)

	a = a + b //алгоритм с суммой
	b = a - b
	a = a - b
	println("a:", a, "b:", b)

	a = a ^ b // a = 01 ^ 10 = 11 = 3 // побитовое исключающее или
	b = a ^ b // b = 11 ^ 10 = 01 = 1
	a = a ^ b // a = 11 ^ 01 = 10 = 2
}