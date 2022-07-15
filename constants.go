package main

import (
	"fmt"
	"math"
)

const s string = "Hello World!!"

func main() {
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Printf("%T", d)
	fmt.Println(int64(d))
	fmt.Printf("%T", d)

	//A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}
