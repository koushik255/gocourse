package main
import (
	"fmt"
)

// hw to manioulate variables outside of thier funciton body
// references variables form outside its body



func main() {

	// sequence := adder()
	// adder function is only called once, now its the return function and since
	// all of the calls are in main the values stays the same , 
// the return function from the adder function retains the value
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// // value of i keeps updating


	// sequence2 := adder()
	// fmt.Println(sequence2())
	// fmt.Println(sequence2())
	
	subtractor:= func() func(int) int {
		countdown := 99

		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()
	/// sp basically you can manipulae varaibles which are like tied to functions,
	// without having to be within said function body

	// using th eclosure subtractor 
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(2))
	fmt.Println(subtractor(3))
	fmt.Println(subtractor(4))
	fmt.Println(subtractor(5))

}

/// we have a function which returns a function
// i is scoped to the adder function, we print the value of the previous value
// this only works because when you call a function to rpint you are going to get
// its return value not its like what it prints.


// adder does not take any arguemtn,
// adder returns a function  it was parameters and return values
func adder() func() int {
	i := 0
	fmt.Println("Previous value of i :",i) // this and aove effect the intial state
	return func() int {
		i++
		return i
	}
}