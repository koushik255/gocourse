package main

import (
	"bufio"
	"fmt"
	"os"
)

// good for large volumes of data
// provides buffering for large inputs
// if your wathcing a video on youtube, when you hit play the movie starts to downloads, the futre of the movie are all getting downloaded to get played as the movie progress if its a 2 hour movie its going to be a big file so its alot of data getting downloaded at one time, if there was no buffering then you would have to download the movie then wtach it
// so buffering allows you to ksip the wait time and download in the background
// it reads and writes data in chuncks
// tranfersing data between chunks
// bufio reader is a striuct which wraps around io reader
// you limit the readstring by using a delimiting character
// useing read method you can limit the amount of bytes the data can input

func main() {
	// wrapp it in the bufio which accepts other readers as input
	// reader := bufio.NewReader(strings.NewReader("Hello, bufio Packageeee!\nHow are you doing?")) // reading source can be text fiels, word file , anything really
	// /// buf.io reader object cna buffer data and can read data in various ways
	// // reader.Read(data) is transfering the data from the bufio instance made with reader into the slice data which can hold 20 bytes

	// //reading the byte slice
	// data := make([]byte, 20)
	// n, err := reader.Read(data) // returns the number of bytes that they read and and error n is the number of bytes that have been read
	// // its reading the data then transferring that data into the byte slice data
	// if err != nil {
	// 	fmt.Println("Error reading:", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes : %s\n", n, data[:n])

	// // readstring reads data up to a certain delimiting character
	// // line stores the line before the newline rune
	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error: reading string:", err)
	// 	return
	// }
	// // since we only read 20 byte of data from the string in read, the data remaing that we have not read is still left in the buffer space and its waiting to be read,
	// fmt.Println("Read string", line)
	// // what is buf.io writineg , struct wraps around a io.writer

	writer := bufio.NewWriter(os.Stdout)

	// wrtiing byte slice
	data := []byte("hello, bufio package!\n")
	n, err := writer.Write(data)
	// this writes the contents of data into the buffer
	if err != nil {
		fmt.Println("Error: writeing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// flushing the buffer to ensure all data is writting to osstdout
	err = writer.Flush() // this just writes any data that is in the buffer
	if err != nil {
		fmt.Println("Error flushing buffer", err)
		return
	}
	// we have to flush the buffer to get the data to show on the screen
	// the writer writes it into the data so how we need to flush it to the underlying writer

	// writing string
	str := "This is a string\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writeing string:", err)
		return
	}

	fmt.Printf("Number of bytes read %d\n", n)
	// flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing the buffer:", err)
		return
	}
}
