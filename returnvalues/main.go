package main
import (
	"fmt"
	"errors"
)


//function can have multiple return values


func main(){

	// func functionName(parameter1 type1, parameter2 type2 .....)(returnType1,returnType2...){
		// code block
		// return returnvalue1,returnvalue2....
	// }

	q,r := divide(10,3)
	fmt.Println("Quotient:", q,"Remainder:", r)

	result, err := compare(3,2)
	if err !=nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func divide(a,b int)(quotient int, remainder int){
	quotient = a/b
	remainder = a%b
	return 
}


// multiple return values is good for error handling
// one of the return values will be an error!


// when you do not have an error to send inplace of the error return value you use nil

func compare(a,b int) (string, error){
	if (a > b) {
		return "A is greater than B!",nil
	} else if b > a {
		return "B is greater than A!",nil
	} else {
		return "", errors.New("Unable to compare which is greater!")
	}
}