/*
	Converts a string to a float
*/

package main

import (
	"fmt"
	"log"
	"strconv"
)

// Simplifies general error handling.
// Might add this one to a package because it would be frequently useful.
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Function for testing the result of using strconv.ParseFloat
func testConv(str string) {
	conv, err := strconv.ParseFloat(str, 64)
	handleError(err)

	fmt.Println("Input is:", str)
	fmt.Printf("Input type is: %T\n", str)
	fmt.Println("Output is:", conv)
	fmt.Printf("Output type is: %T\n\n", conv)
}

func main() {
	// Enter unit tests here:

	testConv("3.141592653598978")
	testConv("-6.72")
	testConv("14")
	testConv("nope")
}
