package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name!!")
	}
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func MultipleHello(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		greeting, err := Hello(name)
		if err != nil {
			log.Fatal("Empty name provided")
		}

		messages[name] = greeting
	}
	return messages, nil
}

func randomGreeting() string {
	greetings := []string{
		"Hi, %v. Welcome",
		"Great to see you Mr. %v",
		"Hail, %v! Well met!",
	}

	return greetings[rand.Intn(len(greetings))]
}
