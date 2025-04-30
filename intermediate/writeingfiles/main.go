package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	defer file.Close()
	// write data to the file
	data := []byte("Hello world!\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writeing to file", err)
		return
	}
	fmt.Println("Data has been written to file succesfully!")

	file1, err := os.Create("strint.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file1.Close()
	// write string data
	str := "Hello this is a string for the write string function"
	_, err = file1.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string to file", err)
		return
	}
	fmt.Println("String has succesfully been written!")

}
