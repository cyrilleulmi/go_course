package section19

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type userEx1 struct {
	First string
	Age   int
}

func ex1() {
	u1 := userEx1{
		First: "James",
		Age:   32,
	}

	u2 := userEx1{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := userEx1{
		First: "M",
		Age:   54,
	}

	users := []userEx1{u1, u2, u3}

	fmt.Println(users)

	marshalledUsers, error := json.Marshal(users)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(marshalledUsers)
}

type ExtendedUser struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func ex2() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)
	var myUser []ExtendedUser
	error := json.Unmarshal([]byte(s), &myUser)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(myUser)
}

func ex3() {
	u1 := ExtendedUser{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In her majesty's royal service",
		},
	}

	u2 := ExtendedUser{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := ExtendedUser{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []ExtendedUser{u1, u2, u3}

	fmt.Println(users)
	json.NewEncoder(os.Stdout).Encode(users)
}

func ex4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

type SortByAge []ExtendedUser

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type SortByName []ExtendedUser

func (a SortByName) Len() int           { return len(a) }
func (a SortByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByName) Less(i, j int) bool { return a[i].Last < a[j].Last }

func ex5() {
	u1 := ExtendedUser{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := ExtendedUser{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := ExtendedUser{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []ExtendedUser{u1, u2, u3}

	println("unsorted")
	fmt.Println(users)
	printUsers(users)

	println("================================")
	println("Sorted by age")
	sortSayings(users)

	sort.Sort(SortByAge(users))
	printUsers(users)

	println("================================")
	println("Sorted by last name")
	sort.Sort(SortByName(users))
	printUsers(users)
}

func sortSayings(users []ExtendedUser) {
	for _, v := range users {
		sort.Strings(v.Sayings)
	}
}

func printUsers(users []ExtendedUser) {
	for _, v := range users {
		printUser(v)
	}
}

func printUser(user ExtendedUser) {
	println("-------------")
	println(user.First)
	println(user.Last)
	println(user.Age)
	println("Sayings")
	for _, v := range user.Sayings {
		println("\t", v)
	}
}
