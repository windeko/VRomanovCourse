package main

import "fmt"

func main() {
	// без значения при инициализации.
	// по-умолчанию значение Int = 0
	var num0 int

	// значение при инициализации
	var num1 int = 1

	// пропуск типа
	var num2 = 20

	// вывод переменных
	fmt.Println(num0, num1, num2)

	// Короткое объявление переменных
	num := 30

	// Инкремент
	num += 1
	// Тоже инкремент
	num++
	fmt.Println(num)

	// Для переменных используется camelCase

	var camelInt = 50
	fmt.Println(camelInt)

	// Объявление нескольких переменных
	var weight, height int = 60, 176

	// Присваивание в уже существующие переменные
	weight, height = 61, 177

	fmt.Println(weight, height)

	// короткое присваивание
	// хотя бы одна переменная должна быть новой!!
	weight, age := 62, 31

	fmt.Println(weight, age)
}
