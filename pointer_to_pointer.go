package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	ptr = &a
	pptr = &ptr
	fmt.Printf("Value of a= %d \n", a)
	fmt.Printf("Value of *ptr= %d \n", *ptr)
	fmt.Printf("Value of **pptr= %d \n", **pptr)
}
