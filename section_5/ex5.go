package sectionfive

import "fmt"

var underlyingType int
var myType myownint

func ExerciseFive() {
	myType = 42
	fmt.Println(myType)
	fmt.Printf("%T\n\n", myType)

	fmt.Println(underlyingType)
	fmt.Printf("%T\n\n", underlyingType)

	underlyingType = int(myType)
	fmt.Println(underlyingType)
	fmt.Printf("%T\n\n", underlyingType)
}
