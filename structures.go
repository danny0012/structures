package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	/*subscriber := []string{}
	subscriber = append(subscriber, "Aman Singh")
	subscriber = append(subscriber, 4.99)
	subscriber = append(subscriber, true)
	fmt.Println(subscriber)*/
	/*subscriber := map[string]float64{}
	subscriber["name"] = "Aman Singh"
	subscriber["rate"] = 4.99
	subscriber["active"] = true
	fmt.Println(subscriber)*/

	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)

	/*var subscriber struct {
		name   string
		rate   float64
		active bool
	}
	subscriber.name = "Aman Singh"
	subscriber.rate = 4.99
	subscriber.active = true
	fmt.Println("Name:", subscriber.name)
	fmt.Println("Monthly rate:", subscriber.rate)
	fmt.Println("Active?", subscriber.active)*/

	/*var pet struct {
		name string
		age  int
	}
	pet.name = "Max"
	pet.age = 5
	fmt.Println("Name:", pet.name)
	fmt.Println("Age:", pet.age)*/

	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)

	var subscriber1 subscriber
	subscriber1.name = "Aman Singh"
	fmt.Println("Name:", subscriber1.name)
	var subscriber2 subscriber
	subscriber2.name = "Beth Ryan"
	fmt.Println("Name:", subscriber2.name)
	
	myCar := car{name: "Corvette", topSpeed: 337}
	fmt.Printf("%v\n", myCar)
}
