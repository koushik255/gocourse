package main
import (
	"fmt"
)


// a defer statement invokes a function whose excution is deffered to the moment
// the surroudning function returns, either because the surrounding fucntion
//  return satement reach end of body or the go routine is paniciy
// go routine are thronw into the back until they finish there work
// not blocking the main thread, then it comes back when its finished



func main() {

	
	process(10)
	
	// for i:=0;i<5;i++{
	// 	fmt.Println(i)
	// }


}

func process(i int){
	defer fmt.Println("Deferred i value:", i )
	defer fmt.Println(" first Deferred statement executed!") // prints last because it is defered
	//to the end of the function call
	defer fmt.Println("Second Deferred statement executed!")
	defer fmt.Println("Third Deferred statement executed!")
	// third goes first, then second, then first. 
	i++
	fmt.Println("Normal print statment execution")
	fmt.Println("ACutal i value:",i)
}



