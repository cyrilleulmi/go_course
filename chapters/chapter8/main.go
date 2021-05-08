package main

import "log"

func main() {
	var whatToSay, _ string = saySomething("Something")
	var hello, world string = saySomething("Hello")
	var number int = 42

	log.Println(whatToSay)
	log.Println(hello + " " + world)
	log.Println(number)
}

func saySomething(stringToReturn string) (string, string) {
	return stringToReturn, "world"
}
