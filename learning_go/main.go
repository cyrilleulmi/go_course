package main

import "log"

func main() {
	var whatToSay string = saySomething("something")
	log.Println(whatToSay)
}

func saySomething(stringToReturn string) string {
	return stringToReturn
}
