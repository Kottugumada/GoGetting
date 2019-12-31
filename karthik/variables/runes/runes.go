package main

import (
	"fmt"
	"reflect"
)

func main() {
	// creating a rune
	rune1 := 'B'
	rune2 := 'g'
	rune3 := '\a'

	// printing rune and its type
	fmt.Printf("Rune 1: %c; Unicode %U; Type: %s", rune1, rune1, reflect.TypeOf(rune1))
	fmt.Printf("Rune 2: %c; Unicode %U; Type: %s", rune2, rune2, reflect.TypeOf(rune2))
	fmt.Printf("Rune 3: %c; Unicode %U; Type: %s", rune3, rune3, reflect.TypeOf(rune3))
}
