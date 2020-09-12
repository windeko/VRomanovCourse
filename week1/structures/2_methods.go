package main

import "fmt"

// Метод - это функция, которая привязана к определенному типу данных

type Person2 struct {
	id		int
	name	string
}

type Acc struct {
	id		int
	name	string
	person	Person2
}

// не изменит оригинальной структуры для которой вызван,
// так как метод получает данные, а не адрес переменной. В даннам случае смысла не имеет.
func (p Person2) UpdateName (name string)  {
	p.name = name
}

// изменит оригинальную структуру.
// * - передает ссылку на область памяти переменной
func (p *Person2) SetName (name string)  {
	p.name = name
}


// Методы могут применяться не только для структур, но и для всех типов данных
// Например создадим тип MySlice - слайс интов
type MySlice []int

// Метод добавления в MySlice
func (s *MySlice) Add(num int)  {
	*s = append(*s, num)
}

// Количество элементов
func (s *MySlice) Count() int {
	return len(*s)
}

func main()  {
	pers := Person2{
		id:   1,
		name: "Vasya",
	}

	pers.UpdateName("Vasilii")
	fmt.Printf("UpdateName: %#v \n", pers)

	pers.SetName("Russian Ivan")
	fmt.Printf("SetName: %#v \n", pers)

	// Метод можно вывать через родительскую структуру
	acc := Acc{
		id:     1,
		name:   "rivan",
		person: pers,
	}
	acc.person.SetName("Rus Ivan")
	fmt.Printf("Acc SetName: %#v \n", acc)

	// Методы для Слайса
	myS := MySlice{1, 2}
	myS.Add(5)
	fmt.Println("MyS count:", myS.Count())
}
