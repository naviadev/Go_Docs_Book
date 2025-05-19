package main

import (
	"fmt"
	"greetings"
)

func main() {
	message, error := greetings.Hello("야호")
	if error != nil {
		fmt.Print(error)
	} else {
		fmt.Print(message)
	}
}
