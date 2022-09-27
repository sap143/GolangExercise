package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	names := []string{"salman", "arun", "sunil", "preeti"}
	// message, err := greetings.Hello("Salman")
	// message, err := greetings.Hello("")
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
