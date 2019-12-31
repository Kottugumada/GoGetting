package main

import (
	"fmt"
)

func main() {
	// creating a seed
	seed := []rune{'¶', 'þ', '©', '§', 'ƒ'}

	// displaying the details in seed one at a time
	// using range loop
	for index, value := range seed {
		fmt.Printf("\n Character: %c, Unicode: %U, Position: %d ", value, value, index)
	}

}
