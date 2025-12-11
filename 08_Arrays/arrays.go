package main

import "fmt"

// numbered sequence of specific length
func main() {
	// zero values in int bydefault
	var nums [4]int

	// Array Length
	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(nums)

	// false value by default
	var vals [4]bool
	// vals[2] = true
	fmt.Println(vals)

	// Empty String
	var name [3]string
	name[0] = "golang"
	fmt.Println(name)

	// To declare it in single line
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	// 2d arrays
	arr2 := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(arr2)

	// - fixed size, that is predictable
	// - Memory Optimization
	// - Constant time access

	arr3 := [3][3][1]int {{{1},{2},{3}},{{4},{5},{6}},{{7},{8},{9}}}
	fmt.Println(arr3)

}
