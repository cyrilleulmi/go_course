package section11

import "fmt"

func ex1() {
	x := [5]int{1, 22, 315, 47, 52}
	fmt.Printf("%T\n", x)

	for _, v := range x {
		fmt.Println(v)
	}
}

func ex2() {
	slice := []int{1, 2, 33, 42, 5, 64, 37, 81, 94, 10}

	for _, v := range slice {
		fmt.Printf("%v\n", v)
	}

	fmt.Printf("%T\n", slice)
}

func ex3() {
	initialSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	firstSlice := initialSlice[:5]
	secondSlice := initialSlice[5:]
	thirdSlice := initialSlice[2:7]
	fourthSlice := initialSlice[1:6]

	fmt.Println(firstSlice)
	fmt.Println(secondSlice)
	fmt.Println(thirdSlice)
	fmt.Println(fourthSlice)
}

func ex4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)

	fmt.Println(x)
}

func ex5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)

	fmt.Println(y)
}

func ex6() {
	x := make([]string, 0, 50)
	fmt.Printf("%p\n", x)
	fmt.Println(cap(x))

	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`,
		`Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		`Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`,
		`Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		`New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		`Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `Utah`,
		` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Printf("%p\n", x)
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func ex7() {
	superslice := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	for _, slice := range superslice {
		for _, v := range slice {
			fmt.Println(v)
		}
	}
}

func ex8To10() {
	m := map[string][]string{}

	m["bond_james"] = []string{"martini", "women", "something"}
	m["m_miss"] = []string{"james", "computers", "dont care"}
	m["no_dr"] = []string{"sharks", "lasers", "1 billion dollars"}

	fmt.Println(m)
	delete(m, "m_miss")

	for key, value := range m {
		fmt.Println("Key: ", key)

		for i, v := range value {
			fmt.Println("\tIndex: ", i, "\tValue:", v)
		}
		fmt.Println("--------------")
	}
}
