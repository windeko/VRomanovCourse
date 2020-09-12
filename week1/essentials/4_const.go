package main

import "fmt"

const pi = 3.14157

const (
	hello = "Привет"
	e     = 2.718
)

// iota - автоинкремент внутри объявления констант
const (
	zero  = iota
	_     // пустая переменная, пропуск iota
	_     // пустая переменная, пропуск iota
	three // = 3
)
const (
	_         = iota             // пропускаем первое значение
	KB uint64 = 1 << (10 * iota) // 1024
	MB                           // 104857
)
const (
	// нетипизированная константа
	year = 2017

	// типизированная константа
	yearTyped int = 2020
)

func main() {
	fmt.Println(zero, three)
	fmt.Println(KB, MB)
	fmt.Println(year, yearTyped)

	var month int32 = 13

	// Так можно потому что для year будет применено динамическое подставление типа
	fmt.Println(month + year)

	// Так нельзя, потому что типы должны быть одинаковые и нужно сначала привести к одному типу!
	//fmt.Println(month + yearTyped)
}
