package sectionfive

import "fmt"

func SomeFunc() {
	a := 42
	b := "James Bond"
	c := true

	fmt.Println(a, b, c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%v %v %v \n", a, b, c)
	fmt.Printf("%d %b %o \n", a, a, a)
}
