package main
import(
	"fmt"
)

// panic stops normal exuction of a function, it unwinds the stack then exectues any defferd function
//this works UP the stack

func main() {

	// panic(interface()) // interface means you  can input any value
	//  as any type as an arguement for this function interface is like any in javascript

	//defer and panic go hand and hand with each other
	// defer statements are executed when panic statements happen
	// panic should be avoided when regular error handling would work insatead
	
	process(8)
	process(-8)

 /// the recover function returns the value passed to the panicing function
 
	
}

func process(input int){

defer fmt.Println("Deferred 1")
defer fmt.Println("Deferred 2")

	if input <0 {
		fmt.Println("Before Panic")
		panic("input must be a non negative number")
		// fmt.Println("After Panic")
		// defer fmt.Println("Defferd 3")
		//anything after panic will not be reached at runtime

	}
	fmt.Println("The number you chose is:",input)
	
}