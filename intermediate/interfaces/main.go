package main

import (
	"fmt"
	"math"
)
// interface is a type, you just have to fufill its requirements to be considered its type, its like a string, you have to have letters or numbers but they must be in double quotes"" 
// discord bot which automatically compresses a video to 8mb or less..
// that could be cold....

// interface provide a set of method signautres
//polymorphisim withouht relieing on explicit inheritance
// interfaces are delcared outside the main function
// if you make the g a G then you could acces it from import main

type geometry interface{
	area() float64
	perim() float64
}
// geomtry is an interface which has 2 methods with return 2 float64 values
// typee will be of this interface if by default it alligns wiht all the constraits within said interface

type rect struct {
	width, height float64
}

type rect1 struct {
	width, height float64
}


type circle struct {
	radius float64
}

// value reciver since we are not direcitng acces the values 
func (r rect) area() float64{
	return r.height * r.width
}

func (r rect1) area() float64{
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}


func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) diameter() float64 {
	return 2 * c.radius
}
// circle has 3 methods associated with it and rectangle only has 2
// the diameter is not printed because its not area or perim and since we
// only printed the interface diamater does not fullfill any role within
// the interface

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
// takes geomertry as an arguemtn and 



func main() {
	r := rect{width:3,height: 4}
	c := circle{radius:5}
	r1 := rect1{width:3,height: 4}
	measure(r)
	measure(c)
	measure(r1)
	
	// names := []string{"Hello","prabir"}
	// fmt.Println(names)
	// for i,v := range names {
	// 	fmt.Println(i,v)
	// }
	
	myPrinter(1,"JOhn?",45.9,false)
	
	
}
func (r rect1) perim() float64{
	return 2* (r.height +r.width)
}
//  you can also use interfaces as a way to accept anytime which is usefull
// for variadic functions
func myPrinter(i ...interface{}) {
	for _,v := range i {
		fmt.Println(v)
	}
	}
