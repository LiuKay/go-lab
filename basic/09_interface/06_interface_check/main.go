package main

import "fmt"

func main() {
	print(1, 2, "test string")
}

func print(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
		case bool:
		default:
			if v, ok := arg.(string); ok {
				val := v
				fmt.Println("it's a string:", val)
			} else {
				//...
				fmt.Println("not a string:", v)
			}
		}
	}
}
