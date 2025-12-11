package main

// struct je che ae bija struct ma b nakhi sakiye aene embedding kevai

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

// order struct
// composition
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func main() {
	newCustomer := customer{
		name:  "john",
		phone: "1234567890",
	}

	newOrder := order{
		id:       "1",
		amount:   30,
		status:   "received",
		customer: newCustomer,
	}

	newOrder2 := order{
		id:     "1",
		amount: 30,
		status: "received",
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
	}

	fmt.Println(newOrder.customer)
	fmt.Println(newOrder2)
}
