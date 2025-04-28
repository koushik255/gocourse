package main

import (
	"errors"
	"fmt"
)

//errors in Go
func sqrt(x float64)(float64,error) {
	if x <0 {
		return 0, errors.New("Math Error: sqaure root of a negative number")
	}
	// compute sqaure root not really doing htat
	return 1,nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: Empty data")
	}
	return nil // everything would be ok if the data has been processed
	
}


func main(){
	
	
	// result, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)
	
	// result1, err1 := sqrt(-16)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// fmt.Println(result1)
	
	// data := []byte{}
	// // if err := process(data);err != nil {
	// err := process(data) // this is basically the meta when it comes to error
	// // in Go
	// if err != nil {
	// 	fmt.Println("error:",err )
	// 	return
	// }
	// fmt.Println(data)	
	
	err1 := eprocess()
	
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	
}

type myError struct {
	message string 
}


// Capital E uses the built in go package which uses the error interface
//it has a single method which is Error with a capital E which returns are string which returns the error
func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}