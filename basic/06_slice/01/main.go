package main

import (
	"fmt"
)

/**
  数组切片：
	能够自动扩容的数组，len():元素的个数，cap():底层数组的大小，让添加元素超过cap()初始值之后，cap 会翻倍

*/
func main() {
	createSliceBaseArray()

	mySlice := make([]int, 5, 10)
	capOfSlice(mySlice)

	mySlice = appendElement(mySlice)

	capOfSlice(mySlice)

	var results []int
	fmt.Println(results)

	fmt.Println("---mySlice2---")
	mySlice2 := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice2)
	fmt.Println(mySlice2[2:4]) // slicing a slice
	fmt.Println(mySlice2[2])   // index access; accessing by index
	fmt.Println("mySlice2"[2]) // index access; accessing by index
}

func appendElement(arr []int) []int {
	return append(arr, 1, 2, 3)
}

func createSliceBaseArray() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 基于数组创建一个slice slice2:= arr[:]
	var mySlice = arr[5:]

	fmt.Println("Elements of arr:")
	for _, v := range arr {
		fmt.Println(v, " ")
	}

	fmt.Println("Element of slice:")
	for _, v := range mySlice {
		fmt.Println(v, " ")
	}
}

func capOfSlice(arr []int) {
	//mySlice :=make([]int,5, 10)

	fmt.Println("len:", len(arr))
	fmt.Println("cap:", cap(arr))
}

func createSliceDirectly() {
	//mySlice :=make([]int, 5) // 创建个数为5的数组切片，元素初始值为0

	//mySlice2 :=make([]int,0, 10) //allocates an underlying array of size 10 and returns a slice of length 0 and capacity 10
	//	that is backed by this underlying array

	//mySlice3 := []int{1, 2, 3, 4, 5}
}
