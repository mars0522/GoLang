package main

import "fmt"

type kitchen struct {
	numOfPlates int
}

type House struct {
	kitchen
	numOfRooms int
}

func main() {

	h := House{kitchen{10}, 3}
	fmt.Println("House h has this many rooms: ", h.numOfRooms)
	fmt.Println("House h has this many plates: ", h.numOfPlates)
	fmt.Println("The kitchen contents of this house are: ", h.kitchen)
	fmt.Println("The House h has this many plates: ", h.kitchen.numOfPlates)

}
