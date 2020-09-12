package main

import "fmt"

type Person struct {
	id		int
	name	string
	phone	int
}

type Account struct {
	id		int
	name	string
	cleaner	func(str string) string
	owner	Person
}

// МЕТОД - смотри 2_methods.go
func (acc Account) clean() {
	fmt.Println(acc.name, "will clean your room rapidly")
}

func main()  {
	// Полное объявление структуры
	var p1 Person = Person{
		id:    1,
		name:  "Vasilii Petrov",
		phone: 79991112323,
	}

	// Так же можно использовать короткую запись
	p2 := Person{
		id:    2,
		name:  "Petr Vasiliev",
		phone: 79992221212,
	}

	// Можно заполнять не все поля структуры
	acc1 := Account{
		id:   1,
		name: "Test Acc",
	}

	// Короткое объявление структуры
	acc1.owner = Person{
		id:    3,
		name:  "Test Person",
		phone: 123123123,
	}

	fmt.Printf("%#v \n", p1)
	fmt.Printf("%#v \n", p2)
	fmt.Printf("%#v \n", acc1)
	acc1.clean()
}
