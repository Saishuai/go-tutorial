package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// create a [name]message map
	messages := make(map[string]string, 0)

	// loop names, fill in map
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			log.Fatal(err)
			return messages, err
		}

		if _, ok := messages[name]; !ok {
			messages[name] = message
		}
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome to first template.",
		"Nice to meet you, %v",
		"Enjoy your journey! %v.",
	}
	return formats[rand.Intn(len(formats))]
}
