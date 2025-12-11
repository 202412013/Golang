package main

import "fmt"


// by value -> Create distinct copy of value
func ChangeNum(num int) {
	num = 5
	fmt.Println("In ChangeNum", num)
}

// Store Address
func NumChange(num *int){
	// Pointer
	*num = 5 // dereference
	fmt.Println("In NumChange -> ", *num)
}

func main() {
	num := 1
	ChangeNum(num)
	fmt.Println("After changeNum in main -> ", num)

	// Reference => value no address pass karsu. Aene reference pass kariyo kevai
	fmt.Println("Memory Address -> ", &num)
	NumChange(&num)
	fmt.Println("Adter NumChange in main -> ", num)

}
