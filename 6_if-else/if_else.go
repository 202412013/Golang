package main

import "fmt"

func main() {
	age := 13

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is teenager")
	} else {
		fmt.Println("Person is a kid")
	}

	var role = "admin"
	var hasPermissions = true

	if role == "admin" || hasPermissions {
		fmt.Println("Yes")
	}

	if role == "user" && hasPermissions {
		fmt.Println("Yes")
	}

	// We can declare a variable inside if construct
	if age := 21; age >= 18 {
		fmt.Println("Person is an Adult", age)
	}

	// go does not have ternary operator, you will have to use normal if else

}
