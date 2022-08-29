package main

import "fmt"

/*
Задание:
Реализовать паттерн «адаптер» на любом примере.
*/

type Food interface {
	Eat()
}
type Drinks interface {
	Drink()
}

type Human struct {
}

func (h *Human) Cook(f Food) {
	fmt.Println("Человек начал готовить...")
	f.Eat()
}

type Pizza struct {
}

func (m *Pizza) Eat() {
	fmt.Println("И съел пиццу")
}

type DrinksAdapter struct {
	DrinksType Drinks
}

func (da *DrinksAdapter) Eat() {
	fmt.Println("Налил в кружку, ")
	da.DrinksType.Drink()
}

type Lemonad struct {
}

func (l *Lemonad) Drink() {
	fmt.Println("Выпил лимонад")

}

func main() {
	human := new(Human) // инициализируем структуру "человек"
	pizza := new(Pizza) // инициализируем структуру "пицца"
	human.Cook(pizza)   // если человек захочет пиццу, то он его просто приготовит и съест

	lemonad := new(Lemonad)                        // если человек захочет лимонад, то он не сможет его съесть,
	adapter := &DrinksAdapter{DrinksType: lemonad} // поэтому был добавлен адаптер, в котором человек наливает лимонад,

	human.Cook(adapter) //а потом уже пьет

}
