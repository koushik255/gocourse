package main
import (
	"fmt"
)

func main(){

	//... Eclipsis this means the function can accept 0 or more arguemtns of that type
	// func functionName(param1 typ1,param2 typ2, param3 ...typ3) returntyp{
	// 	//function body
	// }
	// there is no limit to the amount of parameter you can input into param/typ 3

	sequence,total := sum(1,20,30,40,50,60)
	fmt.Println("Sequence :", sequence, "Total: ", total)
	sequence2,total2 := sum(2,40,36,40,50,60)
	fmt.Println("Sequence :", sequence2, "Total: ", total2)




	numbers := []int{1,2,3,4,5,9}

	sequence3,total3 := sum(3, numbers...)
	fmt.Println("Sequence :", sequence3, "Total: ", total3)




}

func sum(sequence int, nums ...int)(int, int){
	total := 0
	for _,num := range nums{
		total += num
	}
	return sequence, total
}
