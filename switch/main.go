package main
import "fmt"

func main(){
	// fruit := "apple"
	// switch fruit {
	// case "apple":
	// 	fmt.Println("Its an apple")
	// case "banana" :
	// 	fmt.Println("Its a banada!")
	// default :
	// 	fmt.Println("Unknown fruit!")

	// }

	// 	// multple conditions in a switch statements
	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wenesday", "Thrusday", "Friday":
	// 	fmt.Println("it's a weekay :sadface")
	// case "Sunday", "Saturday":
	// 	fmt.Println("ITS THE WEEKEND!")
	// }

	// //EXpressions in switch statements
	// number := 15

	// switch {
	// case number <10:
	// 	fmt.Println("Numer is less than 10")
	// case number >= 10 && number <20:
	// 	fmt.Println("Number is between 10 and 19")
	// default : 
	// 	fmt.Println("IDK BRO")
	// }
	

	// //switch statement does not flall through by defauly

	// num := 2 

	// switch {
	// case num > 1:
	// 	fmt.Println("greater than 1!")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("The number is 2!")
	// default : 
	// 	fmt.Println("Not 2")
	// }

	checkType(10)
	checkType(3.14)
	checkType("Hello")	
	checkType(true)

}

func checkType (x interface{}){
	switch x.(type){
	case int:
		fmt.Println("its a int")
		fallthrough
	case int32:
		fmt.Println("int32")
	case float64:
		fmt.Println("it is a float")
	case string:
		fmt.Println("Its a string")
	default:
		fmt.Println("its a unknown type")
	}

}