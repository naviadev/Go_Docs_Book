package main

import (
	"fmt"
	"greetings"
)

func main() {
	message := greetings.Hello("야호")
	fmt.Print(message)
}
