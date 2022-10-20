/*
	This script includes a variety of useful functions to condense workflow.
*/

package pkg

import (
	"fmt"
	"log"
)

// HandleError simplifies general error handling.
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Input is a Python-like implementation for getting user input in a condensed format.
func Input(str string) string {
	var response string
	fmt.Print(str)

	_, err := fmt.Scanln(&response)
	HandleError(err)

	return response
}
