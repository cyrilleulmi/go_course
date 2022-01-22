package section17

import "fmt"

func ex1() {
	x := "Test string"
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x)

	fmt.Printf("%T\n", &x)
	fmt.Printf("%v\n", &x)
}

func ex2() {
	myPerson := person{name: "noone", age: 23}

	fmt.Println("persons name: ", myPerson.name, "is", myPerson.age, "years old")

	changeMe(&myPerson)

	fmt.Println("persons name: ", myPerson.name, "is", myPerson.age, "years old")
}

type person struct {
	name string
	age  int
}

func changeMe(personToChange *person) {
	personToChange.name += " the"
	(*personToChange).name += " third"
	personToChange.age += 12
}
