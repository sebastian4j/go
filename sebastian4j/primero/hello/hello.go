package main

// usar: /home/sebastian/go/bin/dlv-dap dap y correr el debuger
import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"

	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Go())
	message, err := greetings.Hello("Sebastián")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"Xanxo", "Xanxa", "Osito"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}
