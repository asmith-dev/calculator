/*
	Testing for how Go orders inputs and how Go permits reused values
*/

package main

import "fmt"

func main() {
	// Map being tested, aptly names "stuff"
	stuff := make(map[int]string)

	// Three inputs, intentionally out of order to see how the map stores data
	stuff[2] = "two"
	stuff[7] = "seven"
	stuff[1] = "one"

	// Loop to check map values
	for num, str := range stuff {
		fmt.Println(num, str)
	}

	// Two more inputs to see if map allows multiple keys per value (or vice versa)
	stuff[2] = "not_two"
	stuff[8] = "seven"

	// Loop to recheck map values
	fmt.Println()
	for num, str := range stuff {
		fmt.Println(num, str)
	}

	/*
		In conclusion, Go inputs maps in order of entry, allows for multiple different keys with the same value,
		but overwrites reused keys.
	*/
}
