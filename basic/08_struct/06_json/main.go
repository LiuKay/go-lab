package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//exported
type person struct {
	Name string
	Age  int
}

//unexported
type personNew struct {
	name string
	age  int
}

//tagged
type personTag struct {
	Name string `json:"name"`
	Age  int    `json:"test_age"`
}

func main() {
	p1 := person{
		Name: "kay",
		Age:  18,
	}
	result, _ := json.Marshal(p1)
	fmt.Println(result)
	fmt.Printf("%T \n", result) //[]uint8
	fmt.Println(string(result)) // {"Name":"kay","Age":18}

	p2 := personNew{
		name: "kay",
		age:  18,
	}
	result1, _ := json.Marshal(p2)
	fmt.Println(result1)
	fmt.Printf("%T \n", result1) //[]uint8
	fmt.Println(string(result1)) //{}

	p3 := personTag{
		Name: "kay",
		Age:  18,
	}
	result2, _ := json.Marshal(p3)
	fmt.Println(result2)
	fmt.Printf("%T \n", result2)
	fmt.Println(string(result2)) // {"name":"kay","age":18}

	// Unmarshal
	fmt.Println("------Unmarshal---------")
	var p4 personTag
	fmt.Println(p4.Name, p4.Age)
	bytes := []byte(`{"name":"kay","test_age":18}`)
	json.Unmarshal(bytes, &p4)
	fmt.Println(p4)

	//encoder
	json.NewEncoder(os.Stdout).Encode(p4)

	//decoder
	var p5 person
	rdr := strings.NewReader(`{"Name":"James", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p5)
	fmt.Printf("%T \n", rdr)
	fmt.Println(p5)
}
