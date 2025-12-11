package main

import "fmt"

// generics :- used karsu duplicate code remove karva mate

// any or interface{}
// int | string -> only int and string allow
func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// comparable -> interface che jema comparable types allow che booleans, numbers, strings, pointers, channels, arrays, structs whose fields are all comparable types.
func printSlice2[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// LIFO
type stack[T interface{}] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3}
	names := []string{"golang", "typescript"}
	vals := []bool{true, false, true}
	printSlice(nums)
	printSlice(names)
	// printSlice2
	printSlice2(nums, "A")
	printSlice2(names, "B")
	printSlice2(vals, "C")

	myStack := stack[int]{
		elements: []int{1, 2, 3},
	}

	myStack2 := stack[string]{
		elements: []string{"1", "3", "2"},
	}

	fmt.Println(myStack)
	fmt.Println(myStack2)
}
