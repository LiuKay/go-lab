package main

import (
	"fmt"
	"sort"
)

type student struct {
	Name string
	Age  int
}

type StudentComparator []student

func (s StudentComparator) Len() int {
	return len(s)
}

func (s StudentComparator) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s StudentComparator) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func main() {
	students := []student{
		{"kay", 18},
		{"kb", 24},
		{"kayb", 22},
	}
	fmt.Println(students)
	sort.Sort(StudentComparator(students))
	fmt.Println(students)
}
