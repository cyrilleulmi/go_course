package section7

import "fmt"

const (
	year_a = 2021 + iota
	year_b = 2021 + iota
	year_c = 2021 + iota
	year_d = 2021 + iota
)

func ex4() {
	mynumber := 42
	fmt.Printf("%d %b %x \n\n", mynumber, mynumber, mynumber)

	mynumber = mynumber << 1
	fmt.Printf("%d %b %x \n\n", mynumber, mynumber, mynumber)
}

func ex5() {
	fmt.Println(year_a)
	fmt.Println(year_b)
	fmt.Println(year_c)
	fmt.Println(year_d)
}
