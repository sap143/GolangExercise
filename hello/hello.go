package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// message := greetings.Hello("Salman")
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
