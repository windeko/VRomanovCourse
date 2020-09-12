package main

import "fmt"

// Обычное объявление
func singleIn(in int) int {
	return in
}

// Много параметров
func multIn(a int, b int, c int) int {
	return a + b + c
}

// Именованный результат
func namedReturn() (out int) {
	out = 42
	return
}

// Несколько результатов
func multipleReturn(in int) (int, error) {
	if in > 5 {
		return in, fmt.Errorf(`in > 5`)
	}
	return in, nil
}

// Множественный Именованный результат
func multWithErr(in int) (res int, err error) {
    if in > 5 {
		err = fmt.Errorf(`in > 5`)
	}
	res = in
	return
}

// Вариативные функции - с неограниченным кол-вом параметров
func sum(in ...int) (result int) {
    for _, val := range in {
        result += val
    }
    return
}

func main() {
    mrs, err := multipleReturn(4)
    fmt.Println(mrs, err)
	fmt.Println(namedReturn())
	mwe, err := multWithErr(1)
	fmt.Println(mwe, err)
	fmt.Println(sum(2, 3, 5))
}
