package main

import "fmt"

func main() {
	grades := [3]int{97, 85, 98}
	fmt.Printf("Grades= %v\n", grades)

	// no size declaration
	fmt.Printf("array without size declaration\n")
	grades1 := [...]int{97, 85, 98}
	fmt.Printf("Grades= %v\n", grades1)

	fmt.Printf("empty array declaration\n")
	var students [3]string
	fmt.Printf("Students: %v", students)

	fmt.Printf("access specific index\n")
	students[0] = "Smith"
	fmt.Printf("Students: %v", students)

	fmt.Printf("access specific indices\n")
	students[0] = "Smith"
	students[1] = "Stan"
	students[2] = "Sean"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of Students: %v\n", len(students))

}
