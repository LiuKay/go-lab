## Hello Wolrd

国际惯例，学习一门新语言，第一个 Demo 就来试试 Hello World 程序吧！

用你选好的IDE 新建一个 hello.go 的文件，或者简单点，用文本编辑器输入如下代码并保存：

```golang
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

在当前文件夹打开命令行，运行：
```sh
go run hello.go
```

输出：
hello world

## Go 命令说明

`go build`:
命令可以将Go语言程序代码编译成二进制的可执行文件，但是需要我们手动运行该二进制文件

`go run`:命令则更加方便，它会在编译后直接运行Go语言程序，编译过程中会产生一个临时文件，但不会生成可执行文件，这个特点很适合用来调试程序

`go help`: 帮助文档

其他命令参考官方文档：
[https://golang.org/cmd/go/](https://golang.org/cmd/go/)