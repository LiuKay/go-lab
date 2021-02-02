package main

import (
	"fmt"
	"strconv"
)

func main() {
	intToFloat()
	runeToString()
	bytesToString()

	//	ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, i, u)

	//	FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
	w := strconv.FormatBool(true)
	x := strconv.FormatFloat(3.1415, 'E', -1, 64)
	y := strconv.FormatInt(-42, 16)
	z := strconv.FormatUint(42, 16)

	fmt.Println(w, x, y, z)

}

func intToFloat() {
	var x = 12
	var y = 12.1230123
	fmt.Println(y + float64(x))
	// conversion: int to float64
}

func runeToString() {
	var x rune = 'a' // rune is an alias for int32; normally omitted in this statement
	var y int32 = 'b'
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(string(x))
	fmt.Println(string(y))
}

func bytesToString() {
	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'}))
	fmt.Println([]byte("hello"))
}

func atoi() {
	var x = "12"
	var y = 6
	z, _ := strconv.Atoi(x)
	fmt.Println(y + z)
}
