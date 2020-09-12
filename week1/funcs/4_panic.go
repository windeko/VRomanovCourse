package main

import "fmt"

// panic - абсолютно исключительная ситуация, которая может повесить демон полностью
// Отлавливается Panic с помощью блока defer
// Чтобы восстановиться из Паники - существует оператор recover() который возвращает ошибку или nil

// Никогда не используй Panic как замену try/catch!!

func deferTest2()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic happened:", err)
		}
	}()
	fmt.Println("deferTest Some Useful Work")
	panic("Something bad happend")
	return
}

func main()  {
	deferTest2()
	return
}
