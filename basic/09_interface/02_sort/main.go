package main

import (
	"fmt"
	"sort"
)

type people []string // people is same as sort.StringSlice

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

/**
sort.Sort(sort.Interface)
*/
func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	fmt.Println("---sortStringSlice---")
	sortStringSlice()
}

func sortStringSlice() {
	strings := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(strings)

	//sort.StringSlice(strings).Sort()
	//sort.Strings(strings)
	sort.Sort(sort.StringSlice(strings))
	fmt.Println(strings)

	sort.Sort(sort.Reverse(sort.StringSlice(strings)))
	fmt.Println(strings)
}

func sortIntSlice() {
	ints := []int{3, 1, 8, 33, 2, 12}
	fmt.Println(ints)
	sort.Sort(sort.IntSlice(ints))
	//sort.Ints(ints)
	//sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)
}
