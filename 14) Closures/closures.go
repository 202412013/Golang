package main

import "fmt"

func counter() func() int {
	var count int = 0

	// function ni andar variable j che ae outer scope nu use kare che toh ae variable nu reference store karine rakhi de. Jethi
	// execute thaya pachi b remove(terminate) na kari de.
	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	fmt.Print(increment(), " ")
	fmt.Print(increment(), " ")
	fmt.Print(increment())
}
