package main
import (
	"fmt"
)





func main(){
	// func <name> (parameter list)(returning multiple valies){ context of function}

	

// 	printHello()

// 	fmt.Println(sum(5,5))

// 	//assigning function to a variales/type

// 	operation := sum
// 	result := operation(3,6)
// 	fmt.Println(result)

// }

// func printHello(){
// 	fmt.Println("HELLO!")

// passing a function as an argument
result := applyOperation(5,3, sum)
fmt.Println("5 + 3 =",result)

 // returning and using a function 
 multiplyBy2 := createMultipler(2)
 fmt.Println("6 *2 =", multiplyBy2(6))


}


func sum(x int, y int)(int){
	return x+ y
}
//first class obkjects, are things that have no restrictions on their use, so basically kind of global
//it means you cna preform a wide range of modifications on it just like strings and integers. 


// function that takes a function as a an argument
func applyOperation(x int,y int,operation func(int,int)int ) int {
	return operation(x,y)
}

//Function that returns a function
func createMultipler(factor int) func(int) int {
	return func (x int ) int {
		return x * factor
	}
}
// this means that functions are a first class object within GO
