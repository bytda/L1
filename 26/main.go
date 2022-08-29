package main

import (
	"fmt"
	"strings"
)

/*Задание:
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	var str string //переменная для проверяемой строки
	str = "abcd"
	fmt.Printf("%s - %v \n", str, UniqChars(str))
	str = "abCdefAaf"
	fmt.Printf("%s - %v \n", str, UniqChars(str))
	str = "aabcd"
	fmt.Printf("%s - %v \n", str, UniqChars(str))
}

func UniqChars(str string) bool { //функция проверки на уникальность
	runes := make(map[rune]struct{}) //создаем мапу рун для хранения символов строки
	str = strings.ToLower(str)       //переводим все символы в нижний регистр
	for _, s := range str {          //обходим строку
		if _, ok := runes[s]; ok { //проверяем наличие символа по ключу
			return false //если есть совпадение, то символ не уникальный и выводим false
		}
		runes[s] = struct{}{} //если совпадения нет - заносим символ в мапу
	}
	return true //если совпадений не было выводим true
}
