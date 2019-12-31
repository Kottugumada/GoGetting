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

	fmt.Printf("make function\n")
	a1 := make([]int, 3)
	fmt.Println(a1)
	fmt.Printf("Slice of length= %v\n", len(a1))

	a2 := make([]int, 3, 100)
	fmt.Println(a2)
	fmt.Printf("Slice of length= %v and the underlying array of capacity = %v\n", len(a2), cap(a2))

	fmt.Printf("append capacity\n")
	a3 := []int{}
	fmt.Println(a3)
	fmt.Printf("Slice of length= %v and the underlying array of capacity = %v\n", len(a3), cap(a3))

	a3 = append(a3, 1)
	fmt.Println(a3)
	fmt.Printf("Slice of length= %v and the underlying array of capacity = %v\n", len(a3), cap(a3))
}
