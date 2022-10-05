/*
	Practicing appending to an array (technically a slice).
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Test slice, aptly named "mkay" because everything is gonna be mkay.
	var mkay []string

	// Several appends using for loop, and each update is printed with length
	for i := 0; i < 10; i++ {
		fmt.Println(mkay, len(mkay))
		mkay = append(mkay, strconv.Itoa(i))
	}
	fmt.Println(mkay, len(mkay))
}
