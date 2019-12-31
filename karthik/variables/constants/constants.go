package main

import "fmt"

const myConst int16 = 27

const (
	a = iota
	b
	c
)

const (
	a1 = iota
	b1 = iota
	c1 = iota
)

const (
	_ = iota // one time write only variable
	cat
	dog
	snake
)

const (
	_ = iota + 5 // allows arithmatic operations
	cats
	dogs
	snakes
)

func main() {
	fmt.Printf("%v, %T\n", myConst, myConst)
	// shadowing
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// add constant to a variable
	var myConst1 int = 29
	fmt.Printf("%v, %T\n", myConst+myConst1, myConst)

	// enumerarted constants
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	fmt.Printf("enumerates automatically\n")
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", b1)
	fmt.Printf("%v\n", c1)

	fmt.Printf("_ one time write only\n")
	fmt.Printf("%v\n", cat)

	fmt.Printf("_ allows arithmatic operations on iota\n")
	fmt.Printf("%v\n", cats) // in cases we need a fixed offset
}
