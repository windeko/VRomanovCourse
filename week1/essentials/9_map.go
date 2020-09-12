package main

import "fmt"

func main() {
	// Инициализация при создании
	var user map[string]string = map[string]string{
		"name":     "Vasya",
		"lastName": "Ivanov",
	}

	// Инициализация пустой мапы размерностью 10
	profile := make(map[string]string, 10)

	fmt.Printf("%d %+v\n", user, profile)

	// Если нет ключа - вернет дефолтное значение для типа
	var midName = user["midName"]
	fmt.Println("midName: ", midName)

	// проверка на существование ключа
	midName, midNameExists := user["midName"]
	fmt.Println("midName: ", midName, "midNameExists: ", midNameExists)

	// Пустая переменная, только проверяем ключ
	_, midNameExists = user["midName"]
	fmt.Println("midNameExists: ", midNameExists)

	// Удаление значения из Мапы
	fmt.Println("User: ", user)
	delete(user, "lastName")
	fmt.Println("User: ", user)
}
