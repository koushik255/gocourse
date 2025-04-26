package main
import (
	"fmt"
)
//pointers in go is not that big of a concept tbh you arelready know them
// stores the memory address of another varialble

// var ptr *int -- pointer which points to an integers value
//type of pointer neeeds to match the type of variable its referenceing
// use a * to set the pointer then use a & to set that pointer to a reference then use another //
// * to dereference the potiner and to obtain the value stored at the memory address

// rare real world use cases of pointer to a pointer
func main() {

	var ptr *int 

	var a int = 19
	 ptr = &a // referecning
	fmt.Println(a)
	fmt.Println(&a)


	fmt.Println(ptr)
	fmt.Println(*ptr) // dereferceng
	// if ptr == nil {
	// 	fmt.Println("Pointer is nil!")
	// }
	fmt.Println("New values")
	modifyValue(ptr)
	fmt.Println(a)
	fmt.Println(*ptr)

}

func modifyValue(ptr *int ) {
	*ptr ++
}


//Go does not support potiner arithmetic Pointers are limited to referencing and derefertingn 
// you can acces struct fields aswell
