package main

import "fmt"

// enumerated types

// custom types
// type MyType string

type OrderStatus int

const (
	Received  OrderStatus = iota // iota ek int hoi che je increment thatu jai che etle aa 0 status kevai
	Confirmed                    // 1
	Prepared                     // 2
	Delivered                    // 3

)

type OS string

const (
	received  OS = "received"
	confirmed    = "confirmed"
	prepared     = "prepared"
	delivered    = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func COS(status OS) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Delivered) // ctrl + space
	COS(delivered)
}
