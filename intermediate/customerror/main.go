package main

import (
	"errors"
	"fmt"
)

// customs errors  allows to create error types which provide more context or additional information about something which has gone wrong
func main() {

	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return // you need to return here
	}
	fmt.Println("Operation completed succesfull!")
}

type customError struct {
	code    int
	message string
	er      error // error of type error
}

// error returns the error message implementing Error(_) method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error: %d: %s %v ", e.code, e.message, e.er)
}

// Functions that returns a custom Error
// func doSomething() error {
// 	// usering & to acces the memory address
// 	return &customError{
// 		code:    500,
// 		message: "Something when wrong!",
// 	}
// }

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong!",
			er:      err,
		}

	}
	return nil // return nothing is there is no error
}

func doSomethingElse() error {
	return errors.New(" Internal error")
}
