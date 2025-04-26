package main
import (
	"fmt"
)
type Shape struct {
	Rectangle
}

var fullname = map[int]string {1:"Hello",2:"Koushik"}


var myMap2 = map[string]int {"a": 1, "b": 2}
	


type Rectangle struct {
	length float64
	width float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// // pointer receivers 
// func (r *Rectangle) addLength() {
// 	r.length++
// }

//better mthod with pointer receiver
//make this pointer so we cna modify the exsticing values of this struct
func (r *Rectangle) Scale (factor float64){
	r.length *= factor // shorthand of r.length = r.length * factor
	r.width *= factor // same with this
}


func main() {
	// fmt.Println(sum(5,4))
	
	rect := Rectangle { length: 10, width:  9}
	area := rect.Area()
	fmt.Println(area)
	// fmt.Println("Length before:",rect.length)
	// rect.addLength()
	// fmt.Println("Length after", rect.length)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of the rectangle when * by a factor of 2 is now:",area)

	num := MyInt(-5)
	num1 := MyInt(9)
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num.welcomeMessage())
	fmt.Println(fullname)
	
	s := Shape{Rectangle: Rectangle{length: 10,width: 9}}
	// s gets promoted so you can directly access area
	
	fmt.Println(s.Area())
	s.Scale(3)
	area1 := s.Area()
	
	fmt.Println("New area for s after factor of 3",area1)
	
}

	type MyInt int 
	
	// Method on a user defined type
	/// before the function name we need to mention the type
	// like >><<< right there because we need to create an instance to access
	// the value within this type or the MyInt type to be more specific
	
	func (m MyInt) IsPositive() bool{
		return m >0 
	} 
	
	// you can just associate the function so that you make it a 
	// method of that type
	// you only need to call an instance if you are 
	// extracting (trying to print or compare a number) or 
	// Modifying a value(when you would have to use pointer receivers)
	 
	func (MyInt) welcomeMessage() string {
		return "Welcome to MyInt type!"
	}



// func sum(a int, b int) int {
// 	return a + b
// }