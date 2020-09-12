package person

import "fmt"

// Публичные переменные, типы и структуры объявляются с большой буквы
// Доступны всему приложению
var Public = 1

// Приватные переменные, типы и структуры - с маленькой
// Приватные доступны только внутри данного пакета
var private = 2

type Person struct {
	ID		int
	Name	string
	secret	string
}

func (p Person) UpdateSecret(secret string)  {
	p.secret = secret
}

func (p Person) GetSecret() string {
	return p.secret
}

func (p Person) PrintSecret() {
	fmt.Println(p.secret)
}