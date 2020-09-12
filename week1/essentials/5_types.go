package main

import "fmt"

type UserID int

func main() {
	idx := 1
	var uid UserID = 42

	// Даже если базовые типы одинаковые, компилятором они все равно считаются разными типами
	// Код ниже приведет к ошибке типов данных
	//sum := idx + uid

	// Приведение переменной с типом Int (idx) к типу UserId
	myID := UserID(idx)

	fmt.Println(myID, uid)
}
