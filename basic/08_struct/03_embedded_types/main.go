package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	licenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		licenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		licenseToKill: false,
	}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.licenseToKill)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.licenseToKill)
}
