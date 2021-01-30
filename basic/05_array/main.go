package main

import "fmt"

/**
1.Go 中数组是值类型,作为参数输入时会被复制一份，方法内不会修改原始数组
2. 数组的长度在定义之后无法再次修改
*/

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	modify(arr)
	fmt.Println("In main(), array values:", arr)

	iterate(arr)
}

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify(),array values:", array)
}

func iterate(array [5]int) {
	fmt.Println("iterate:")
	for i, v := range array {
		fmt.Println("Array element:[", i, "]=", v)
	}
}
