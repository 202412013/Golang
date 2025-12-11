package main

import "fmt"

// Variadic functions (rest operator in JS)
// nums slice array bani jase
// Khali int value joise
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// koi b type ni value chalse
func print(nums ...interface{}) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func main() {
	// Println bau badha parameter lai sake che
	fmt.Println(1, 2, 3, 4, 5, "Hi", "bro")

	result := sum(1, 2, 3, 4, 5)
	// slice array already che and sum karavu che pn tya already slice array banave che toh spread operator use karsu JS na jm
	nums := []int{1,2,3,4,5}
	fmt.Println("Spread Operator sum -> ",sum(nums...))  // spread Operator

	fmt.Println("Rest Operator sum -> ", result)

	// print karse aa fucntion because aapde logic ae rite lakhiyu che
	print(3, "Hi", true)

}
