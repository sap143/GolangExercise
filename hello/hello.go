package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Salman")
	fmt.Println(message)
}
