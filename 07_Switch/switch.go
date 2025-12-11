package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	i := 5

	// There is no need to write break statement.

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("Other")
		}
	}

	whoAmI("golang")
	whoAmI(333)
	whoAmI(0.333)

}
