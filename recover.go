package main

import "fmt"

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recover form ", r)
	}
}
func fullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name can not be nil")
	}
	if lastName == nil {
		panic("runtime error:  last name can not be nil")
	}
	fmt.Printf("%s %s \n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
func main() {
	defer fmt.Println("deffered call in the main")
	firstName := "Varun"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

}
