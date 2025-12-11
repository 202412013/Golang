package main

import (
	"fmt"
	"time"
)

// class nathi golang ma but aena badle aapde use karsu struct
// OOP use kari sakiye che struct na help thi
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

// Behavioral add kari sakai struct na andar
func (o *order) changeStatus(status string) {
	o.status = status
}

// agar aapde koi b struct ni value ne change na kariye toh jarur nathi pointer ni
func (o order) getAmount() float32 {
	return o.amount
}

// constructor jevu struct ma
func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

func main() {
	// if you don't set any field, default value is zero value
	// int => 0, float => 0, string => "", bool => false
	// amuk vastu add nai kariye toh b problem nai aave (error nai aave)
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received", // createdAt ma {0 0 <nil>}  aay jase
	}

	myOrder.createdAt = time.Now()
	fmt.Println("Order Struct 1 -> ", myOrder)

	// koi field ne get karvu hoi toh kari sakiye che
	fmt.Println(myOrder.status)

	myOrder2 := order{
		id:        "2",
		amount:    100,
		status:    "delivered",
		createdAt: time.Now(),
	}

	fmt.Println("Order Struct 1 -> ", myOrder2)

	myOrder.changeStatus("Confirmed")
	fmt.Println(myOrder.getAmount())

	// constructor call kariye emj function call karvanu
	myOrder3 := newOrder("1", 33.333, "received")
	fmt.Println(myOrder3)

	// Ekj vaar use karvu che struct toh aa rite. struct ne name aapvani b jarur nathi
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

}
