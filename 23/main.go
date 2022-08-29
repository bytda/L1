package main

import "fmt"

/*
Задание:
Удалить i-ый элемент из слайса.
*/
func main() {
	data := []int{3, 1, 4, 7, 2, 55, 1, 9, 5, 3}
	fmt.Println(data)
	i := 4

	data = DirtyDelete(data, i)
	fmt.Println("Грязный результат: ", data)

	data = CleanDelete(data, i)
	fmt.Println("Чистый результат: ", data)
}

func DirtyDelete(data []int, i int) []int { // способ с нарушением порядка (но с фиксированной скоростью)
	fmt.Println("Грязный метод: ", data)
	data[i] = data[len(data)-1] //копируем последний элемент слайса на позицию которую надо удалить
	//получается на заданной позиции и в конце одинаковые значения
	data = data[:len(data)-1] //урезаем слайс с конца на одну позиции тем самым удаляя значение которое мы копировали
	return data               //возращаем
}

func CleanDelete(data []int, i int) []int { // чистый способ (но с линейной скоростью)
	fmt.Println("Чистый метод: ", data)
	copy(data[i:], data[i+1:]) //выбираем отрезок слайса после элемента который нам нужно удалить
	//и вствляем этот отрезок  в отрезок начиная с элемента удаления, тем самым нужный элемент удаляется,
	// а в конце среза остается нулевой обьект
	data = data[:len(data)-1] //урезаем срез с конца на одну позицию
	return data               //возращаем
}
