package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println(v1)
	fmt.Println("Length of the map: ", len(m))

	delete(m, "k2")
	fmt.Println("Now the map length: ", len(m))

	_, prs := m["k2"]
	fmt.Println("Prs:", prs)

	n := map[string]int{"foo": 1, "shoo": 2}
	fmt.Println("map: ", n)

}
