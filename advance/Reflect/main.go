package main

import (
	"fmt"
	"reflect"
)

func main() {

	x := 42
	v := reflect.ValueOf(x)
	t := v.Type()

	fmt.Println("Value :", v)
	fmt.Println("Type :", t)
	fmt.Println("Kind :", t.Kind()) // kind returns the kind of the type

	fmt.Println("Is zero:", v.IsZero())
}
