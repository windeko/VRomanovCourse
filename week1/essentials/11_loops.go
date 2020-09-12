package main

import "fmt"

func main() {
	// Цикл без условия, while(true) OR for(;;;)
	for {
		fmt.Println(`Loop iteration`)
		break
	}

	// Цикл с условием, while(isRun)
	isRun := true
	for isRun {
		fmt.Println("Loop with condition")
		isRun = false
	}

	// Цикл с условием и блоком инициализации
	for i := 0; i <= 3; i++ {
		if i == 1 {
			// Служебное слово, переход к следующей итерации цикла!!!
			continue
		}
		fmt.Println("Loop with Continue", i)
	}

	// Операции по Slice
	s1 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(s1); i++ {
		fmt.Println("Slice iteration:", s1[i])
	}

	// OR BETTER!!!
	for idx, val := range s1 {
		fmt.Println("Element with index", idx, "has value =", val)
	}

	// Операции по Map
	profile := map[string]string{"name": "Vasya", "midName": "Aleksandrovich"}

	for key, val := range profile {
		fmt.Println(key, ":", val)
	}

	// Итерация по строке
	str := "Привет, Мир!"
	for pos, char := range str {
		fmt.Printf("position: %d, char: %#U\n", pos, char)
	}
}
