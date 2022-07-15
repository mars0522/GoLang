package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}

	sum := 0

	for i, num := range nums {
		sum += num
		fmt.Println("i:", i)
	}
	fmt.Println("sum of the array: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Printf("Key:", k)
	}

	for i, c := range "go" {
		println(i, c)
	}
}
