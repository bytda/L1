package main

import (
	"fmt"
	"math/big"
)

/*
Задание:
Разработать программу, которая перемножает, делит,
складывает, вычитает две числовых переменных a, b, значение которых > 2^20.
*/

func main() {
	BigMath()
}

func BigMath() {
	fmt.Println("BigMath started")
	a := new(big.Int)                                     //объвляем большую переменную
	a.SetString("18000000000000000000000000000000", 10)   // 18 * 10^30 в качестве строки пишем значение , после запятой десятичность
	b := new(big.Int)                                     //объвляем большую переменную
	b.SetString("5600000000000000000000000000000000", 10) // 56 * 10^32
	fmt.Println("a: ", a)                                 //выводим
	fmt.Println("b: ", b)

	sum := new(big.Int)
	sum.Add(a, b) //сумма
	fmt.Println("a+b: ", sum)

	dif := new(big.Int)
	dif.Sub(a, b) //вычитание
	fmt.Println("a-b: ", dif)

	mult := new(big.Int)
	mult.Mul(a, b) //умножение
	fmt.Println("a*b: ", mult)

	div := new(big.Int)
	div.Div(b, a) //деление
	fmt.Println("b/a: ", div)

}
