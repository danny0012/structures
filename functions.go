package main

import "fmt"

type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts)

	p := minimumOrder("Hex bolts")
	fmt.Println(p.description, p.count)

	subscriber1 := defaultSubscriber("Aman Singh")
	applyDiscount(subscriber1)
	printInfo(subscriber1)
	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)
	
	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)
	var pointer *subscriber = &s
	fmt.Println((*pointer).rate)
	fmt.Println(pointer.rate)
}
