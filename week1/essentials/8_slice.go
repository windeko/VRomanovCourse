package main

import "fmt"

func main() {
	var buf0 []int
	buf1 := []int{}
	buf2 := []int{42}
	buf3 := make([]int, 0)
	buf4 := make([]int, 5)
	buf5 := make([]int, 5, 10)

	fmt.Println(buf0, buf1, buf2, buf3, buf4, buf5)

	// Обращение к элементу слайса
	fmt.Println(buf5[2])

	// Добавление элеементов в слайс
	buf5 = append(buf5, 6, 7, 8, 9)
	fmt.Println(buf5)

	// Добавление другого слйса
	buf6 := make([]int, 3)
	buf6 = append(buf6, buf2...)
	fmt.Println(buf6)

	// просмотр информации о слайсе:
	// len() - длина элементов
	// cap() - вместимость слайса
	var length, capacity = len(buf5), cap(buf5)
	fmt.Println(length, capacity)

	// СРЕЗЫ
	buf := []int{1, 2, 3, 4, 5, 6, 7}

	sli1 := buf[1:4] // 2,3,4
	sli2 := buf[:2]  // 1,2
	sli3 := buf[4:]  // 5,6,7

	fmt.Println(sli1, sli2, sli3)
	fmt.Println("=========")

	// Изначально sli3 и sli4 смотрят на одну и ту же область памяти
	sli4 := sli3[:]

	fmt.Println(sli3, sli4)

	// поэтому изменение элементов одного - приводят к изменениям в другом
	sli4[0] = 20

	fmt.Println(sli3, sli4)

	// А вот тут происходит переполнение Capacity и соответственно выделение новой области памяти для sli4
	sli4 = append(sli4, 10)

	fmt.Println(sli3, sli4)

	// Теперь изменения в sli4 не ведут к изменениям в sli3
	sli4[0] = 1

	fmt.Println(sli3, sli4)
	fmt.Println("=========")

	// КОПИРОВАНИЕ ОДНОГО СЛАЙСА В ДРУГОЙ
	// Создаем пустой слайс с размерностью и вместимостью слайсо который мы собираемся копировать
	sli5 := make([]int, len(sli4), len(sli4))

	// и копируем
	copy(sli5, sli4)

	fmt.Println(sli4, sli5)

	// Так как это полноценные копии - изменения в одном слайсе не ведут к изменениям в другом
	sli4[0] = 0
	sli5[1] = 42

	fmt.Println(sli4, sli5)
	fmt.Println("=========")

}
