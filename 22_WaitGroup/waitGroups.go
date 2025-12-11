package main

import (
	"fmt"
	"sync"
)

// waitgroup pointer hovu joie
func task(id int, w *sync.WaitGroup) {
	defer w.Done() // defer -> function execute thaya pachi last ma run thai che. Similar to useEffect hook cleaning feature.
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // counter type mechanism
		go task(i, &wg)
	}

	wg.Wait()
	// WaitGroup ma 3 vastu yaad rakhvani ->
	// 1) waitgroup add karo
	// 2) done karo
	// 3) wait karo
}
