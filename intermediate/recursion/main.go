package main
import (
	"fmt"
)

// where a function calls itself directly or indirectly in order
// to solve a problem
// breaks something down into smaller problems until those problems become
// simple enough to where it can solve it
// Recusion depends on your base case



func main() {

	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(12))
	fmt.Println(sumOfDigits(5))
	fmt.Println(sumOfDigits(12345))
}


func factorial (n int) int {
	//base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive case: factorial of n is n * facortial (n-1)
	return  n * factorial( n-1 )
	// n * (n -1) * (n-2) * factorial (n-3) ... basically goes on and on until it 
	// reaches factorial of 0
	// if it reaches 1 then it gives the output of 1
}

func sumOfDigits(n int) int {
	// base case
	if n <10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}