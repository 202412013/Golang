package main

// Race Condition thi bachva mate use kariye che mutex
import (
	"fmt"
	"sync"
)

// Race condition etle ke multiple processes same resources ne modify kare che same time pr aenej race condition che.

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{
		views: 0,
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(myPost.views)
}
