package main

import (
	"fmt"

	"github.com/gophercoders/simpleio"
)

func main() {
	var strangersName string
	var strangersAge int

	fmt.Println("The hello stranger program shows you how to use variables ")
	fmt.Println("to read input from the keyboard.")
	fmt.Println("")

	fmt.Println("Please type in your name.")
	strangersName = simpleio.ReadStringFromKeyboard()
	fmt.Println("Please tell me you age.")
	strangersAge = simpleio.ReadNumberFromKeyboard()
	fmt.Print("Hello ")
	fmt.Println(strangersName)
	fmt.Print("You are ")
	fmt.Print(strangersAge)
	fmt.Println(" years old.")
}
