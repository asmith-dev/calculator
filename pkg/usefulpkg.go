/*
	This script includes a variety of useful functions to condense workflow.
	It only needs added to the same directory as the file using the functions.
	Alternatively, individual functions can be copied and pasted into the file.
	In that case, the "import" statement will need updated accordingly.
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
