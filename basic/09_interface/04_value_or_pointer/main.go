package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle1 struct {
	radius float64
}

//value
func (c circle1) area() float64 {
	return math.Pi * c.radius * c.radius
}

type circle2 struct {
	r float64
}

//pointer
func (c *circle2) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c1 := circle1{radius: 2}
	printArea(c1)
	printArea(&c1) // 编译通过，使用也没问题，因为编译器会根据值 receiver 的值类型自动生成一个带指针的方法实现

	c2 := circle2{r: 2}
	//circle2 does not implement shape (area method has pointer receiver)
	//printArea(c2) // 编译不通过，因为 circle2 实现的 area 方法的 receiver 是一个指针，必须传入一个指针
	printArea(&c2)
}

func printArea(s shape) {
	fmt.Println(s.area())
}
