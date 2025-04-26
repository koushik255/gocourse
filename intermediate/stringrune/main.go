package main
import (
	"fmt"
)
// strings and runes in GO
// string is a sequence of bytes u8 int values which often represent text
// strigngs are inmutable so once they are create they cannot be chaged


func main() {

	message := "Hello, Go!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!"
	rawMessage := `Hello\nGo` // using backticks for raw stinrg literals
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message variable:",len(message))
	
	// byte is unsigned interger 8
	fmt.Println("The first character of message var is:",message[0])
	// this prints th eascii value of the letter H

	greeting := "Hello"
	name := "Alice"
	msg := greeting +" "+ name
	fmt.Println(msg)
}