package main

import "fmt"

func deferTest() string {
	fmt.Println("Defer test execution")
	return "Defer test result"
}

func main()  {
	// defer - ключевое слово, показывающее что эту функцию нужно выполнить по завершению функции main в данном случае
	defer fmt.Println("After work")

	// При чем defer выполняются в обратном порядке, то есть сначала выполнится Second а потом After
	defer fmt.Println("Second defer")

	// Чтобы выполнить некую функцию отложенно - ее нужно обернуть в анонимную функцию
	defer func() {
		fmt.Println(deferTest())
	}()

	fmt.Println("Some useful work")
}
