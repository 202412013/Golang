package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object(Javascript), dictionary(python)
func main() {
	// creating map
	m := make(map[string]string) // map[key]value

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"], m["area"])

	// IMP :- If key value does not exists in map then it returns zero for int, empty for string, false for bool
	m1 := make(map[string]int)
	m1["age"] = 33
	m1["temp"] = 30
	fmt.Println(m1["phone"]) // return 0
	fmt.Println("Length -> ", len(m1))

	// delete temp
	delete(m1, "temp")
	fmt.Println("After delete -> ", m1)

	// Clear Map
	clear(m1)
	fmt.Println("Clear -> ", m1)

	// Shorthand Map creation
	m2 := map[string]int{"price": 51, "phone": 3}
	fmt.Println(m2)

	// v -> Value , ok -> bool (exists or not)
	v, ok := m2["price"]
	fmt.Println(v)
	if ok {
		fmt.Println("All Okay")
	} else {
		fmt.Println("Not Okay")
	}

	m3 := map[string]int{"price": 40, "phones": 3}
	m4 := map[string]int{"price": 40, "phones": 3}

	// m3 and m4 are objects toh aapde == thi check na kari sakiye
	fmt.Println("Maps(m3,m4) are equal or not -> ", maps.Equal(m3, m4))

}
