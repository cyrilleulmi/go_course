package sectionfive

import "fmt"

var x int
var y string
var z bool

func PrintPackageScopeStuff() {
	fmt.Println(x, y, z)
	fmt.Println("Zero Values")
}
