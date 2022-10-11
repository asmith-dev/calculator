/*
	Explores appending to a multidimensional slice
*/

package main

import "fmt"

func main() {
	// The slice and a dummy string
	var multiDim [][][]string
	str := "stuff"

	// Appends to the slice
	multiDim = append(multiDim, [][]string{{str}})
	multiDim = append(multiDim, [][]string{{str}})
	multiDim[0] = append(multiDim[0], []string{str})
	multiDim[0][0] = append(multiDim[0][0], str)

	// Prints to demonstrate the results
	fmt.Println(multiDim)
	fmt.Println(len(multiDim))
	fmt.Println(len(multiDim[0]))
	fmt.Println(len(multiDim[1]))
	fmt.Println(len(multiDim[0][0]))
	fmt.Println(len(multiDim[0][1]))
	fmt.Println(multiDim[0][0][0])
}
