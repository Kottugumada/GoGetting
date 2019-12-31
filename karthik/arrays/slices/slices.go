package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // size need not be defined
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	// slices are naturally reference types
	b := a
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("Slice syntax\n")
	c := a[:]   // slice of all elements
	d := a[3:]  // slice from the 4th element to the end
	e := a[:6]  // slice of the first 6 elements
	f := a[3:6] // slice the 4th, 5th and 6th elements
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
