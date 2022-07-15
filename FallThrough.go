package main

import "fmt"

func main() {

	i := 5
	switch {

	case i <= 10:
		fmt.Println("1st statement executed")
		fallthrough
	case i >= 20:
		fmt.Println("2nd statement executed")
	default:
		fmt.Println("No statement execute")
	}
}
