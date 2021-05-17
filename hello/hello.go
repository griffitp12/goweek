package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Pete")
	fmt.Println(message)
}
