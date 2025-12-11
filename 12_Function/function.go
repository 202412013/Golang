package main

import "fmt"

// return int
func add(a int, b int) int {
	return a + b
}

// return int and datatype same hoi toh last ma je aapiye ae pelana badhama lagi jase
func sum(a, b int) int {
	return a + b
}

// Multiple values return
func getLanguages() (string, string, string, bool) {
	return "golang", "Javascript", "CPP", true
}

// fn khali name che je aapde aapiye che nested function nu name
func processIt(fn func(a int) int) int {
	return fn(3)
}

// Return function
func processfunc() func(a int) int {
	return func(a int) int {
		return a
	}
}

func main() {
	result := add(3, 5)
	fmt.Println(result)

	fmt.Println(sum(6, 3))

	fmt.Println(getLanguages())

	// Destructuring and _ -> print nathi karavu and error na aape etle use krsu underscore
	lang1, lang2, lang3, _ := getLanguages()
	fmt.Println("Destructuring Languages -> ", lang1, lang2, lang3)

	// pass karsu function and return karse value
	fn := func(a int) int {
		return a
	}
	ans := processIt(fn)
	fmt.Println("Nested Function returns int value -> :- ", ans)

	// Function return
	fun := processfunc()
	ans2 := fun(3)
	fmt.Println("Function return value -> ", ans2)

}
