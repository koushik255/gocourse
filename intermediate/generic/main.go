package main

import "fmt"

// func swap[T any](a,b T) (T,T ){
// 	return b,a
// }
// function that just swaps 2 values around

type Stack[T any] struct {
	elements []T
}

type fullName struct {
	firstName string
	lastName string
	
	
}
// t can be anytype /interface()
// t is a slice which holds elements of type T so it can hold any value
// but they have to be the same type not any mixed types


func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements,element)
	// adding element to the existing elements
	// and we use pointer so we can change the slice in the Stack struct
}

func (s *Stack[T] )pop() (T,bool){
	if len(s.elements) == 0{
		var zero T
		
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
	// we return the element that we are deleting and we return true because we have succesfully deleted the last element so we are basically just confirming that we did indeed just delete the last element from the slice

}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}
// checks if the slice has any elements inside of it
// value reciver since not modify values
func (s Stack[T]) printAll() {
	if len(s.elements) == 0{
		fmt.Println("The stack is emppty")
		return
	}
	fmt.Println("Stack elements: ")
	for _,v := range s.elements {
		fmt.Println(v)
	}
	fmt.Println()
}


func main (){
	// x,y := 1,2
	// x,y = swap(x,y)
	// fmt.Println(x,y)
	
	// x1,y1 := "John","Koushik"
	// x1,y1 = swap(x1,y1)
	// fmt.Println(x1,y1)
	
	
	
	shik := fullName {firstName: "Koushik", lastName: "smedga"}
	fmt.Println(shik)
	
	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	fmt.Println(intStack.pop()) // returns what its deleing and returns if it was succesful or not
	intStack.printAll() 
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty?",intStack.isEmpty())
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty?",intStack.isEmpty())
	
	// using same struct but for string
	stringStack := Stack[string]{}
	stringStack.push("hello")
	stringStack.push("world")
	stringStack.push("john")
	stringStack.printAll()
	stringStack.pop()
	stringStack.printAll()
	stringStack.pop()
	stringStack.pop()
	fmt.Println("Is the stack empty:",stringStack.isEmpty())
	
	
}