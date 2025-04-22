package main
import (
	"fmt"
)


// range provides a convinient way to interate over data structures

func main(){

	

	// numbers := [5]int{1,2,3,4,5,}
	// for i,v :=range numbers {
	// 	fmt.Println(i,v)
	// }

	// sum := 0
	// for i:=0;i < 5; i++{
	// 	sum = sum +1
	// 	fmt.Println(sum)
	// }


	message := "Hello world"

	for i,v := range message{
		fmt.Println(i,v)
		fmt.Printf("Index: %d,Rune: %c\n",i,v)
	}
	//for arrays slices and strings the range iterates in order from th efirst
	//element to the last element
	//for maps range iterates over key value pairs but idk nobody knows really how 
	//it chooses to iterate
	
}

