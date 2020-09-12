package main

import (
	"./person"
	"fmt"
)

func main()  {
	p := person.NewPerson(1, "test", "test")
	fmt.Println(*p)
	fmt.Println(p.GetSecret())
}
