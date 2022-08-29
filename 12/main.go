package main

import "fmt"

/*
Задание:
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

type MapSet struct { //объвляем интерфейс
	data map[string]struct{}
}

func main() {
	items := []string{"cat", "cat", "dog", "cat", "tree"} // объвляем слайс переменных

	var s MapSet //объвляем обьект интерфейса
	fmt.Println("Последовательность строк:", items)
	s.Insert(items...)
	fmt.Println("Созданное множество: ", s.GetAll())
	fmt.Println("Есть ли элемент \"sun\"? ", s.Check("sun"))
	fmt.Println("Есть ли элемент \"dog\"? ", s.Check("dog"))
	s.Delete("dog")
	fmt.Println("Есть ли элемент \"dog\" после удаления? ", s.Check("dog"))

}

func (s *MapSet) Insert(strs ...string) { // Insert добавляет в множество новые элементы (или не добавляет если они там уже есть)
	for _, str := range strs { //перебираем множество
		if s.data == nil { //проверяем проинициализирована ли мапа
			s.data = make(map[string]struct{}) //если нет - создаем
		}
		s.data[str] = struct{}{} // если элемент уже был добавлен, то ничего не изменится,
		// если не был, то он добавится (особенность использования map в основе множества)
	}
}
func (s *MapSet) Check(str string) bool { //функция проверки наличия элемента
	_, ok := s.data[str] //проверяем наличие такого ключа
	return ok            //возращаем булевую переменную
}
func (s *MapSet) Delete(str string) { //объвялем функцию удаления
	if s.Check(str) { //проверяем если существует запрошенная позиция
		delete(s.data, str) // если да - удаляем
	}
}
func (s *MapSet) GetAll() []string { //функция возращающая все элементы
	var str []string           //создаем слайс для вывода
	for item := range s.data { //обходим мапу
		str = append(str, item) //записыаем элементы в слайс
	}
	return str //возраащем слайс

}
