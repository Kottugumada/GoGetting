package main

import "fmt"

func main() {
	var n bool
	fmt.Println(n)
	fmt.Printf("%v %T\n", n, n)

	// compare
	m := 1 == 1
	p := 1 == 2
	fmt.Printf("%v %T\n", m, m)
	fmt.Printf("%v %T\n", p, p)

	// arithmatic operations
	fmt.Printf("arithmatic operations\n")
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// bit operators
	// a 1010
	// b 0011
	fmt.Printf("bit operations\n")
	fmt.Println(a & b)  // 0010 AND
	fmt.Println(a | b)  // 1011 OR
	fmt.Println(a ^ b)  // 1001 XOR
	fmt.Println(a &^ b) // 0100 AND NOT

	// bit shift
	fmt.Printf("bit shift\n")
	c := 8              // 2^3
	fmt.Println(c << 4) // 2^3 * 2^4 = 2^7
	fmt.Println(c >> 4) // 2^3 / 2^4 = 2^-1

	// type conversion
	var a1 int = 20
	var b1 int8 = 5
	fmt.Println(a1 + int(b1))

}
