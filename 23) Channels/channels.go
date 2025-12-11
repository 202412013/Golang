package main

import (
	"fmt"
	"time"
)

// Channels -> Pipe jevu hoi che means from one side we are sending data and from another side we are receiving data.
// Multiple goroutines work thai rayu che and data pass karvo che 2 goroutines ma toh aena mate use kariye che Channels
// Deadlock -> Ek process run ho rahi hai aur dusri process wait kar rahi hai lekin peli process release nai kar pa rahi hai.

// Sending Data
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing number", num)
		time.Sleep(time.Second)
	}
}

// Receiving Data
func sum(result chan int, num1 int, num2 int) {
	ans := num1 + num2
	result <- ans
}

// waitgroup na help thi je kaam kari sakiye ae same kaam aapde channel na help thi b kari sakiye che
// goroutine synchronizer
func task(done chan bool) {
	defer func() {
		done <- true
	}()
	fmt.Println("processing...")
	//done <- true
}

// Buffer Channel
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() {
		done <- true
	}()

	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("Done Sending...")

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)

	// This is Important
	close(emailChan)
	<-done

	// messageChan := make(chan string)
	// messageChan <- "ping"  // blocking
	// msg := <- messageChan
	// fmt.Println(msg)

	// numChan := make(chan int)
	// go processNum(numChan)

	// numChan <- 5

	// // Infinite
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result // blocking

	// fmt.Println(res)

	// done := make(chan bool)
	// go task(done)

	// <-done // block

}
