package main

import "fmt"

var v1 int = 100

func main() {
	// fmt.Println("Hello")
	// fmt.Println(42)
	// var keyword, name and type
	var i int
	i = 27
	fmt.Println(i)

	var j int = 28
	fmt.Println(j)

	// don't tell what datatype the variable is
	p := 29
	fmt.Println(p)
	// %v gives me the value
	// %T gives me the type
	fmt.Printf("%v, %T ", p, p)

	fmt.Println()

	var q float32 = 28
	fmt.Printf("%v, %T ", q, q)

	// shadowing scope
	fmt.Println()

	fmt.Println(v1)
	var v1 int = 101
	fmt.Println(v1)

	// var block
}
