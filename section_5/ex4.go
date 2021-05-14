package sectionfive

import "fmt"

type myownint int

var myOwnVariable myownint

func PrintOwnType() {
	fmt.Println(myOwnVariable)

	myOwnVariable = 42
	fmt.Printf("%T\n", myOwnVariable)
	fmt.Println(myOwnVariable)
}
