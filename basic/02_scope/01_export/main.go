package main

import (
	"fmt"
)

func main() {
	TestFunc()
	fmt.Println(ExportName)

	// 编译错误,无法引用未导出的属性：cannot refer to unexported name inner_package.inName
	//fmt.Println(inner_package.inName)

	//cannot refer to unexported name inner_package.inFunc
	//inner_package.inFunc()
}
