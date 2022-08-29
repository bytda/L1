package main

import "fmt"

func main() {
	// создаем слайс
	data := []int{0, 5, 7, 4, 7, 4, 11, 54, 23, 11, 13, 12, 35, 48, 79, 45, 8, 46}
	fmt.Println(data)

	// сортируем слайс
	QSort(data, 0, len(data)-1)
	fmt.Println(data)
}

// применяем разбиение Хоара
func QSort(data []int, low, high int) {
	if low < high { //если длина массива меньше 2, то сортировать его нет смысла
		var p int
		oporElement := data[low+(high-low)/2] // выбираем опорный элемент как среднюю позицию в массиве
		// (в меньшую строну если в массиве четное кол-во элементов)
		i := low  //первый индекс указываем как начало отрезка сортировки
		j := high // второй индекс указываем как конец отрезка сортировки
		for {
			for data[i] < oporElement { // движемся слева и ищем элемент больший либо равный опорному
				i++
			}
			for data[j] > oporElement { // движемся справа и ищем элемент меньший либо равный опорному
				j--
			}
			if i >= j { // если индексы встретились, выходим из цикла перебора элементов
				p = j
				break
			}
			data[i], data[j] = data[j], data[i] // если не встретились, меняем местами данные и сдвигаем индексы на один шаг
			i++
			j--
		}

		QSort(data, low, p) // рекурсивно вызываем функцию
		QSort(data, p+1, high)
	}

}