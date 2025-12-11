package main

import "fmt"

const s = "Hi"
const global int = 333

// := will not allow
// name:="Will give error for expected declaration"

func main() {
	// infer
	const s = "Abc" // will print

	fmt.Println(s, global)

	// Multiple Constants
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)

}
