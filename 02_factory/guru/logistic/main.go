package main

import "fmt"

func main() {
	p := &Package{
		Quanity:     1000,
		Description: "Something",
		OtherInfo:   "Rest Info",
	}
	fmt.Println("Choose sea logistic")
	logistic := Logistic{
		kind: "sea"}
	logistic.planDelivery()
	sts, err := logistic.vehicle.Deliver(p)
	fmt.Printf("Deliver status %d and error: %v\n", sts, err)
}
