package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	b := a
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println()
	// manipulate a value in b
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// point to the location
	fmt.Println()
	b1 := &a
	fmt.Println(a)
	fmt.Println(b1)
}
