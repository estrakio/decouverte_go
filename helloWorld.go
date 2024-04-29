package main

import (
	"fmt"
	"log"
	"test_1/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Batard")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
