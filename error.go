package main

import (
	"fmt"
	"os"
)

func main() {
	panic("a problme")

	_, err := os.Create("/temp/file")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
