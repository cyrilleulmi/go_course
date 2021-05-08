package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to ", myString)

	changeUsingPointer(&myString)
	log.Println("after func call: ", myString)

}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
