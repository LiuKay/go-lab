package main

import "fmt"

// 大写开头的属性或方法将会对外暴露
var ExportName = "kay"

// 小写开头的属性或方法对外不可见
var inName = "Liu"

func TestFunc() {
	fmt.Println("test func")
}

func inFunc() {
	fmt.Println("inner func")
}
