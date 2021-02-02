package main

import (
	"fmt"
	"math"
)

// shape 是一个接口
type shape interface {
	area() float64
}

// square 是一个实现类
type square struct {
	side float64
}

//方法实现
func (s square) area() float64 {
	return s.side * s.side
}

// 另一个实现类
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var s shape = square{side: 5} // 声明一个接口 shape 类型，并赋值为实现类 square
	fmt.Printf("%T\n", s)         // 打印类型
	info(s)
	accept(s)

	c := circle{radius: 2}
	info(c)
	accept(c)
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func accept(any interface{}) {
	fmt.Printf("Accept any type:%T\n", any)
}
