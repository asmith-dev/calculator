/*
	Demonstrates the usage of strings.ContainsAny
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Strings of chars ordered by types
	ops := "+-*/"                        // operators
	nums := "1234567890"                 //numbers
	caps := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // capital letters
	lows := "abcdefghijklmnopqrstuvwxyz" // lowercase letters

	// A test string and a slice to log the types of each char in the test string
	test := "Ab6+lD-*0"
	var typesOfTest []string

	// Loop iterates through the test string
	for i := 0; i < len(test); i++ {
		temp := string(test[i])

		// Categorizes each char in the test string
		switch {
		case strings.ContainsAny(ops, temp):
			typesOfTest = append(typesOfTest, "OPS:"+temp)
		case strings.ContainsAny(nums, temp):
			typesOfTest = append(typesOfTest, "NUMS:"+temp)
		case strings.ContainsAny(caps, temp):
			typesOfTest = append(typesOfTest, "CAPS:"+temp)
		case strings.ContainsAny(lows, temp):
			typesOfTest = append(typesOfTest, "LOWS:"+temp)
		}
	}

	// Prints the results
	fmt.Println(typesOfTest)
}
