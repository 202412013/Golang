package main

// goroutines na help thi aapde CPU incentive kari sakiye concurrently, background workers kari sakiye, etc

import (
	"fmt"
	"time"
)

// Goroutines -> golang nu one of the most powerful features che goroutines
// Goroutines lightweight thread che jyare aapde multithreading karvi hoi, concurrently run krvu hoi toh ae aapde krsu goroutines na help thi

func task(id int) {
	fmt.Println("doing task", id)
}

func main() {

	// Fast che blocking karta -> goroutines

	for i := 0; i < 10; i++ {
		go task(i) // goroutines feature and parallel(concurrent) run krse.

		// Inline goroutines
		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)
	}

	time.Sleep(time.Second * 2) // main function ne rokine rakhiyu che thodi vaar mate
}
