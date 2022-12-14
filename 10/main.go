package main

import "fmt"

/*
Задание:
Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	data := [...]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} //задаем массив
	group := make(map[int][]float32)                                        //создаем мапу срезов для группировки
	for _, da := range data {                                               //перебираем элементы массива
		g := (int(da) / 10) * 10 //преобразуя в int дробная часть отпадаем,
		//делением и умножением на 10 получаем разраяд десяток числа
		group[g] = append(group[g], da) //записываем по ключу в мапу числа
	}
	for i, float32s := range group { //перебираем мапу
		fmt.Printf("%v : %.1f \n", i, float32s) //выводим значения
	}
}
