package main

import "fmt"

func main() {
	order := &LambSkewer{}
	orderAdd := &Fat{skewer: order}
	orderAdd2 := &LemonSlices{orderAdd}
	fmt.Printf("The price is %v\n", orderAdd2.getPrice())
	fmt.Printf("Your oder is %v", orderAdd2.getInfo())
}
