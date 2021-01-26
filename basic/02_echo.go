package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	echoCommandLineArgs()
}

func echo1() {
	// 注意 string 的默认是是 “” 而不是 java中的 null
	var s, sep string
	// for 循环的写法
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//// 类似于 while
	//count := 1
	//for count<10 {
	//	fmt.Print(count+ ',')
	//	count++
	//}
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echoCommandLineArgs() {
	for index, value := range os.Args[1:] {
		fmt.Println(strconv.Itoa(index) + ":" + value)
	}
}
