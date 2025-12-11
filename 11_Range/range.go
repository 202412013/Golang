package main

import "fmt"

// Iterating over data structures
func main() {
	nums := []int{1, 2, 3}

	// for loop and print in next line
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i], " ")
	}
	// for loop and print in one line
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}

	fmt.Println() // next line

	// sum with the help of Range in Slice
	sum := 0
	// i -> index and num -> value
	for i, num := range nums {
		sum = sum + num
		fmt.Println("Index and value => ", i, num)
	}
	fmt.Println(sum)

	// Range with Map
	m := map[string]string{"fname": "john", "lname": "doe"}

	for index, value := range m {
		fmt.Println("Map -> ", index, value)
	}

	for index := range m {
		fmt.Print("Single value gives you -> ", index)
		fmt.Println(" && Value access by single key -> ", m[index])
	}

	// range in string (unicode code point rune)
	// i -> starting byte of rune
	// Less than 255 -> store karse 1 byte maj
	// More than 255 -> take more one byte
	for i, c := range "golang" {
		fmt.Println(i, string(c), c)
	}

}
