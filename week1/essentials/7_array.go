package main

import "fmt"

func main() {
	// Размер массива является частью его типа

	// инициализация массива значениями по-умолчанию
	var a1 [3]int // [0,0,0]
	fmt.Println("short: ", a1)
	fmt.Printf("short %v\n", a1)
	fmt.Printf("full %#v\n", a1)

	// определение размеера массива при инициализации
	a2 := [...]int{1, 2, 3}
	fmt.Println(a2)

	// Выходить за пределы массива нельзя!!
	//a2[3] = 12
}
