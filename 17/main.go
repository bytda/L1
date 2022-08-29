package main

import (
	"errors"
	"fmt"
)

/*
Задание:
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	data := []int{0, 1, 2, 4, 6, 7, 9, 9, 12, 13, 15, 16, 17, 25, 25, 26} // создаем слайс для поиска

	idx, err := BinarSearch(data, 16) // вызываем функцию возвращающую индекс искомого числа (либо ошибку)

	if err != nil { // если возникла ошибка выводим ее

		fmt.Println(err)
		return
	}
	// если ошибки нет, то выводим индекс числа
	fmt.Println("Индекс числа:", idx)
}
func BinarSearch(data []int, value int) (int, error) {
	if len(data) < 1 { //проверяем чтобы срез не был пуст
		err := errors.New("Массив пуст")
		return 0, err
	}
	low := 0              //начальная точка поиска
	high := len(data) - 1 //конечная точка поиска

	for { //обходим срез
		mediana := low + (high-low)/2

		if data[mediana] == value { //если в найденном индексе хранится искомое значение,

			return mediana, nil // возвращаем индекс
		}

		if high-low == 0 { // если отрезок поиска имеет нулевую длину (один элемент),
			err := errors.New("Значение не найдено") //то выводим ошибку об отсутствии искомого значения
			return 0, err
		}
		if data[mediana] < value { // если центральное значение меньше искомого,
			low = mediana + 1 //то сдвигаем левую границу отрезка поиска
			continue
		}

		if data[mediana] > value { // если центральное значение больше искомого,
			high = mediana - 1 //то сдвигаем правую границу отрезка поиска

		}
	}

}
