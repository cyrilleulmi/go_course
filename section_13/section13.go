package section13

import "fmt"

type person struct {
	first            string
	last             string
	favoriteIceCream []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func ex1() {
	p1 := person{
		first:            "obi",
		last:             "wan",
		favoriteIceCream: []string{"vanilla", "burnt"},
	}

	p2 := person{
		first:            "anakin",
		last:             "skywalker",
		favoriteIceCream: []string{"chocolate", "not sand"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for _, v := range p1.favoriteIceCream {
		fmt.Println(v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for _, v := range p2.favoriteIceCream {
		fmt.Println(v)
	}
}

func ex2() {
	p1 := person{
		first:            "obi",
		last:             "wan",
		favoriteIceCream: []string{"vanilla", "burnt"},
	}

	p2 := person{
		first:            "anakin",
		last:             "skywalker",
		favoriteIceCream: []string{"chocolate", "not sand"},
	}

	peopleMap := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for key, val := range peopleMap {
		fmt.Println(key)
		fmt.Println(val.first)
		fmt.Println("favorite Ice Creams")
		for _, v := range val.favoriteIceCream {
			fmt.Println("\t", v)
		}
		fmt.Println("----------------")
	}
}

func ex3() {
	myTruck := truck{
		vehicle: vehicle{
			doors: 3,
			color: "yellow",
		},
		fourWheel: true,
	}

	mySedan := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println(myTruck)
	fmt.Println("doors: ", myTruck.doors)
	fmt.Println("fourWheel: ", myTruck.fourWheel)

	fmt.Println(mySedan)
	fmt.Println("doors: ", mySedan.doors)
	fmt.Println("luxury: ", mySedan.luxury)
}

func ex4() {
	myAnonymousStruct := struct {
		name string
		age  int
	}{
		name: "dan",
		age:  12,
	}

	fmt.Println(myAnonymousStruct)
}
