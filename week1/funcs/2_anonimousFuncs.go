package main

import (
	"fmt"
	"strconv"
)

// Обычная функция
func doNothing() {
	fmt.Println(`I'm a regular function'`)
}

func main()  {
	// Анонимная функция
	func (in string) {
		fmt.Println(`anon func out:`, in)
	}("Hello world!")

	// Присвоение анонимной функции в переменную
	myFunc := func(in string) {
		fmt.Println(`anon through var says:`, in)
	}
	myFunc("Moo")

	// Функция с коллбэком - полезно когда по завершении функции нужно вызвать другую функцию!!
	timesTwo := func(in int, callback func(string)) {
		times := in * 2
		callback(strconv.Itoa(times))
	}
	timesTwo(21, myFunc)

	// Замыкания
	prefixer := func(prefix string) func(str string) {
		return func(str string) {
			fmt.Printf("[%s] %s \n", prefix, str)
		}
	}
	successLog := prefixer("SUCCESS")
	errorLog := prefixer("ERROR")
	successLog("Good Job!!")
	errorLog("Try it one more time...")
}
