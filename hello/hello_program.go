package main

import (
	"log"

	"example/greetings"
)

func main() {
	log.SetPrefix("Message: ")
	log.SetFlags(0)
	// message, error := greetings.Hello("Naveen")
	// if error != nil {
	// 	log.Fatal("Emplty name provided")
	// }
	// log.Print(message)
	// if error != nil {
	// 	log.Print(error)
	// }
	names := []string{
		"Naveen",
		"Rajesh",
		"Ummesh",
	}
	messageMap, err := greetings.MultipleHello(names)
	if err != nil {
		log.Fatal("Something wrong")
	}

	log.Print(messageMap)
}
