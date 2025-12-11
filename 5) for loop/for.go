package main

import "fmt"

// for ekluj che looping ma golang ni andar.
// while nathi golang ma

// aapde while loop banaie che for loop thi.

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// Infinite loop
	// for {
	// 	println("1")
	// }

	fmt.Println("Classic For Loop")

	// classic for loop
	// In Go, else cannot be used directly with continue or break
	for i := 0; i <= 3; i++ {
		if i == 2 {
			continue
		} else if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// 1.22 version new feature :- range
	fmt.Println("Range new Feature in version 1.22 :- ")
	for i := range 3 {
		fmt.Println(i)
	}

}
