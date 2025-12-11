package main

import (
	"fmt"
	"slices"
)

// slice dynamic array (Work like a vector)
// most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nil(null) and we will not give size then golang will understand that it is slice array
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println("Length -> ", len(nums))

	// Not nil and by default 0
	// Arguments :- arr, length (capacity if capacity not given), elements to be insert (capacity)
	var arr = make([]int, 3, 5)

	fmt.Println(arr == nil)
	fmt.Println(arr)
	// Capacity -> maximum numbers of elements can fit
	fmt.Println("capacity -> ", cap(arr))

	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	fmt.Println(arr)
	fmt.Println("Capacity -> ", cap(arr)) // capacity 10 because we are pushing 6 element. (Same like vector)

	fmt.Println("Best way to create Slice Array")
	// Best way to initialized slice array
	var arr2 = make([]int, 0, 3)
	arr2 = append(arr2, 1)
	arr2 = append(arr2, 2)
	arr2 = append(arr2, 3)
	fmt.Println(arr2)
	fmt.Println("Capacity -> ", cap(arr2))

	fmt.Println("Shorthand to create Slice Array")
	// Shorthand slice array creation
	arr3 := []int{} // {} -> define empty Slice
	fmt.Println(arr3)
	fmt.Println(arr3 == nil) // Empty Slice not a nil Slice
	fmt.Println("Length of array before -> ", len(arr3))
	arr3 = append(arr3, 1)
	fmt.Println(arr3)
	fmt.Println("Capacity -> ", cap(arr3))
	fmt.Println("Length of array -> ", len(arr3))

	var arr4 = make([]int, 0, 3)
	arr4 = append(arr4, 1)
	var arr5 = make([]int, len(arr4)) // minimum length 1 hovi joie toj copy thase

	// Copy Function
	copy(arr5, arr4)
	fmt.Println(arr4, arr5)

	// Slice Operator
	var n = []int{1, 2, 3}
	fmt.Println(n[0:2]) // return [1 2]
	fmt.Println(n[:1])  // return [1]
	fmt.Println(n[1:])  // start from index 1 and return [2 3]

	// slice package
	var n1 = []int{1, 2, 3}
	var n2 = []int{1, 2, 3}

	fmt.Println(slices.Equal(n1, n2))  // return type bool
	fmt.Println(slices.Concat(n1, n2)) // bau badha functions che slices na andar

	// 2d Array
	var n3 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(n3)

}
