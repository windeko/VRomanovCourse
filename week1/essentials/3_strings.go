package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// по-умолчанию строка пустая
	var str string

	// со спец символами
	var hello string = "Привет\n\t"

	// без спец символов
	var world string = `Мир\n\t`

	// UTF-8 из коробки
	var helloWorld = "Привет, Мир!"

	// Одинарные кавычки для байт (uint8)
	var rawBinary byte = '\x27'

	// Конкатенация строк
	andGoodMorning := helloWorld + " и доброе утро!"

	fmt.Println(str, hello, world, helloWorld, rawBinary, andGoodMorning)

	// Получение длинны строки
	byteLen := len(andGoodMorning)
	symLen := utf8.RuneCountInString(andGoodMorning)

	fmt.Println(byteLen, symLen)

	// получение подстроки - в байтах, не в символах!!
	helloBytes := helloWorld[:12]
	h := helloWorld[0]
	hString := string(h)

	fmt.Println(helloBytes, h, hString)
}
