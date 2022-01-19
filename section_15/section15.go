package section15

import (
	"fmt"
	"math"
)

func ex1() {
	i := ex_1_foo()
	fmt.Println(i)

	s := ex_1_bar()
	fmt.Println(s)
}

func ex_1_bar() string {
	return "test"
}

func ex_1_foo() int {
	return 42
}

func ex2() {
	xi := []int{1, 2, 3, 4, 562, 34, 24, 213, 41, 12}
	fmt.Println(sumWithVariadic(xi...))
	fmt.Println(sumWithArray(xi))
}

func sumWithVariadic(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return sum
}

func sumWithArray(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return sum
}

func ex3() {
	defer fmt.Println("First line")
	fmt.Println("Second line")
	defer fmt.Println("third line")
}

func ex4() {
	someone := person{
		first: "some",
		last:  "one",
		age:   24,
	}

	someone.speak()
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("hi, I am", p.first, p.last)
}

func ex5() {
	mySquare := square{side: 10}
	myCircle := circle{radius: 10}

	info(mySquare)
	info(myCircle)
}

type square struct {
	side float32
}

type circle struct {
	radius float32
}

type shape interface {
	area() float32
}

func (s square) area() float32 {
	return s.side * s.side
}

func (c circle) area() float32 {
	return c.radius * c.radius * math.Pi
}

func info(s shape) {
	println("the shape has an area of", s.area())
}

func ex6() {
	println(func() string {
		return "hello there"
	}())
}

func ex7() {
	myFunc := func() string {
		return "general kenobi!"
	}

	println(myFunc())
}

func ex8() {
	myFuncOfAFunc := funcOfAFunc()
	println(myFuncOfAFunc())
}

func funcOfAFunc() func() string {
	return func() string { return "func inside func" }
}

func ex9() {
	evenOrOdd(1, evenCallback, oddCallback)
	evenOrOdd(2, evenCallback, oddCallback)
	evenOrOdd(3, evenCallback, oddCallback)
	evenOrOdd(4, evenCallback, oddCallback)
}

func evenCallback() {
	println("even!")
}

func oddCallback() {
	println("odd!")
}

func evenOrOdd(number int, evenCallback func(), oddCallback func()) {
	if number%2 == 0 {
		evenCallback()
	} else {
		oddCallback()
	}
}

func ex10() {
	{
		i := 20
		println(i)
	}
	{
		i := 40
		println(i)
	}
}
