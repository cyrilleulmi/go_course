package sectionfive

import "fmt"

var x2 int = 42
var y2 string = "James Bond"
var z2 bool = true

func SprintStuff() {
	s := fmt.Sprintf("%v\t%v\t%v", x2, y2, z2)
	fmt.Println(s)
}
