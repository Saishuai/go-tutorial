package main

import (
	"example/greetings"
	"fmt"
	"log"

	"rsc.io/quote"
)

func hello_world() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Opt())
}

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"saishuai.yuan",
		"test you",
		// "",
	}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
