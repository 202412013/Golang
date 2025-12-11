package main

import "fmt"

func main() {
	var name string = "golang"

	// infer
	var name2 = "Automatically String"
	var is_Adult = true

	// int automatically assign int32, int64 as per the computer.
	var age int = 10
	// var age int8
	// var age int16
	// var age int32
	// var age int64

	// Shorthand Syntax
	id := 11

	// Value I want to assign later
	var n string
	n = "Value"

	// Float have two types only float32 and float64
	var price float32 = 3.3
	// Another way to make it short = var price = 3.3
	// One more another way and also short way :- price := 333.33

	fmt.Println(name, name2, is_Adult, age, id, n, price)
}
