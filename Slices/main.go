package main
import (
	"fmt"
	"slices"
)

// Slices are references to underlying array
//they do not store any data themselves


func main(){

// 	// var SliceName[]ElementType

// var numbers []int
// var numbers1 = []int{1,2,3} //declared and not used


// numbers2 := []int{9,8,7}  ///without using the var keyword

// slice := make([]int,5)

a := [5]int{1,2,3,4,5}
slice1 := a[1:4]


//Apending values into a slice
slice1 = append(slice1, 6,7)
fmt.Println("Slice 1 original",slice1)


//Copying a slice
sliceCopy := make ([]int,len(slice1))
copy(sliceCopy,slice1)
fmt.Println("Slice copy: ",sliceCopy)

var nilSlice []int 
fmt.Println(nilSlice)



//iteracing over slices

	for i,v := range slice1{
		fmt.Println("Index",i,"value",v)
	}

	fmt.Println("Accesing element at index 3 :",slice1[3])
	



	// slice1[3] = 50
	// fmt.Println("Accesing element at index 3 :",slice1[3])


	if slices.Equal(slice1,sliceCopy){
		fmt.Println("The slices are equal!")
	} else {
		fmt.Println("Slices are not equal")
	}

	//multiple dimentontal slice
	twoD := make([][]int,3)
	for i := 0; i<3; i++ {
		innerLen := i +1
		twoD[i] = make([]int,innerLen)
		for j := 0; j < innerLen ; j++ {
			twoD[i][j] = i + j

		}
	}
	fmt.Println(twoD)


	//slice(low:high) [] not ()
	slice2 := slice1[2:4]
	fmt.Println(slice2)

	fmt.Println("The capacity of slice2 is ", cap(slice2))
	fmt.Println("the len of slice2 is ", len(slice2))
	//once a slice is init it is always associated with a array, a slice is a referecning to 
	//the underlying array and therfore it shares storage with the same array, distinct arrays represents 
	//distinct storage  the capacity is the sum of the length of the arrya and th earray beyond the slice
	//so a slice is just a part of the array
	
	
	Numbers := []int{1,2,3,4,5}

	for i,v := range Numbers {
		fmt.Println(i,v)
	}

	Numbers = append(Numbers, 6,7,8)
	fmt.Println(Numbers)

	for i,v := range Numbers {
		fmt.Println(i,v)
	}

}
