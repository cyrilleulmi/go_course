package section15

import "fmt"

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
