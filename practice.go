package main

import "fmt"

func main() {

	a := map[int]string{

		1: "a",
		2: "b",
	}
	for b, c := range a {
		fmt.Println(b, c)
	}
}
