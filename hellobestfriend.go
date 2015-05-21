package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader
	var myName string
	var bestFriendsName string

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("The hello best friend program prints your name and your")
	fmt.Println("best friends name to the screen.")
	fmt.Println("Both names have to be typed into the program. The program")
	fmt.Println("will store both names in variables.")
	fmt.Println("")
	fmt.Println("Please type in your name.")
	myName, _ = reader.ReadString('\n')
	fmt.Println("Please type in your best friends name.")
	bestFriendsName, _ = reader.ReadString('\n')
	fmt.Print("Hello ")
	fmt.Println(myName)
	fmt.Print("You best friends is, ")
	fmt.Println(bestFriendsName)
}
