package section7

import (
	"fmt"
	"time"
)

func ex1() {
	fmt.Println("12")

	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}
}

func ex2() {
	for i := 65; i < 65+26; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%U '%c\n'", i, i)
		}
	}
}

func ex3() {
	for i := 1995; i <= time.Now().Year(); {
		fmt.Println(i)
		i++
	}
}

func ex4() {
	i := 1995
	for {
		if i > time.Now().Year() {
			break
		}
		fmt.Println(i)
		i++
	}
}

func ex8() {
	switch {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
}

func ex9() {
	favSwitch("something")
	favSwitch("squash")
}

func favSwitch(favSport string) {
	switch favSport {
	case "tennis":
		fmt.Println("tennis")
	case "floorball":
		fmt.Println("floorball")
	case "squash":
		fmt.Println("squash")
	default:
		fmt.Println("football")
	}
}
