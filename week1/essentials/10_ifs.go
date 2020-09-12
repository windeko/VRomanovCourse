package main

import (
	"fmt"
)

func main() {
	// В GO - в блоке условия обязано быть только Булево выражение!! Никаких Int или String!
	boolVal := true
	if boolVal {
		fmt.Println("boolVal is true")
	}

	mapVal := map[string]string{"name": "Vasya", "midName": "Aleksandrovich"}

	// Условие с блоком инициализации
	if keyValue, keyExist := mapVal["name"]; keyExist {
		fmt.Println("name =", keyValue)
	}

	// Получаем только признак существования ключа
	if _, keyExist := mapVal["name"]; keyExist {
		fmt.Println(`variable mapVal["name"] exists`)
	}

	// Множественный if/else
	cond := 1
	if cond == 1 {
		fmt.Println(`Cond = 1`)
	} else if cond <= 5000000 {
		fmt.Println(`Cond <= 5000000!! Cond =`, cond)
	} else {
		fmt.Println(`!Cond =`, cond)
	}

	// Switch по 1 переменной
	switch cond {
	case 1:
		fmt.Println(`Cond = 1`)
		fallthrough // Проваливание в следующее условие, чтобы не произошло выхода после выполнения кейса!!
	case 500:
		fmt.Println(`Cond = 500`)
	case 1000:
	default:
		fmt.Println(`!Cond =`, cond)
	}

	// Switch как замена многим if/else
	val1, val2 := 10, 15
	switch {
	case val1 > 5 || val2 < 10:
		fmt.Println("Nice!")
	case val2 == 15:
		fmt.Println("Keep it up!")
	default:
		fmt.Println("Try one more time")
	}

	// Выход из Switch и выход из Цикла
Loop:
	for key, val := range mapVal {
		fmt.Println("switch in loop", key, val)
		switch {
		case key == "name":
			fmt.Println("Name:", val)
			break
			fmt.Println("Dumbass: true") // не выведется, потому что после break
		case key == "midName":
			fmt.Println("midName:", val)
			break Loop // Выход из цикла с меткой Loop
		}
	}

}
