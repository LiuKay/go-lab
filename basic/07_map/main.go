package main

import "fmt"

func main() {
	fmt.Println("---nil---")
	var myMap map[string]string
	fmt.Println(myMap)
	fmt.Println(myMap == nil)

	fmt.Println("---init---")
	var person = map[string]interface{}{}
	// other format:
	// person := make(map[string]interface{})
	// person := map[string]interface{}{}
	// myGreeting2 :=map[string]string{
	//		"Tim":"Good morning",
	//		"Jenney":"Bojour",  // note:","
	//	}

	person["first_name"] = "kay" // add element
	person["last_name"] = "liu"  // add element
	person["age"] = 18

	fmt.Println("Person:", person)
	fmt.Println("len(Person)", len(person)) //len()

	delete(person, "age") // delete key

	//delete(person,"non_exist") // delete not exist key will not error

	// 取值和判断是否存在指定 key
	if val, exists := person["age"]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(person)

	//loop
	for key, value := range person {
		fmt.Println(key, ":", value)
	}

}
