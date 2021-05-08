package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"

	log.Println(myMap["dog"])

	var mySlice []string
	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "Bob")

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 12, 14}
	log.Println(numbers[0:2])
	log.Println(numbers[1:8])
}
