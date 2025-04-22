package main
import "fmt"

//Good way to manage data of the same data type
//lenght of an array of go is fixed and given to the compiler at compile time

func main(){

	// // var arrayName [size]element type

	// var number [5] int
	// fmt.Println(number)	


	// number[4] = 20
	// fmt.Println(number)

	// number[0] = 9
	// fmt.Println(number)

	// //init a array with hold va
	
	// fruits := [4]string{"Apples", "Bananas", "Oranges", "Grapes"}
	// fmt.Println("Fruits array:",fruits)

	// fmt.Println("Third element:", fruits[2])

	// originalArray := [3]int{1,2,3}
	// copiedArray :=originalArray

	// copiedArray[0] = 100

	// fmt.Println("Original Array:", originalArray)
	// fmt.Println("CopiedArray:", copiedArray)

	// for i := 0; i <len(number); i++{
	// 	fmt.Println("Element at index,",i, ":", number[i])
	// }

	// ///modern index enumeratio
	// for i,v := range number {
	// 	fmt.Printf("index: %d, Value: %d\n",i,v)
	// }

	// //underscore is a blank identifer, used to store unused values

	// // a,_ := someFunction()
	// // fmt.Println(a)
	// // // fmt.Println(b)


	// // b := 2
	// // _ = b

	// fmt.Println("the length of the number array is ", len(number))

	// //comparing Arrays
	// array1 := [3]int{1,2,3}
	// array2 := [3]int{1,2,3}
	// fmt.Println(array1 == array2)

	// if (array1 == array2) {
	// 	fmt.Println("Both of the arrays are equal!")
	// } else {
	// 	fmt.Println("The ararys are NOT equal!")
	// }

	// //muktipl dimensional array

	// var matrix [3][3]int = [3][3]int{
	// 	{1,2,3},
	// 	{4,5,6},
	// 	{7,8,9},
	// }
	// fmt.Println(matrix)


	originalArray := [3]int{1,2,3}
	var copiedArray *[3]int
	//..this is a pointer which points to an array of 3 integers


	copiedArray = &originalArray

	copiedArray[0] = 100

	fmt.Println("Original Array:", originalArray)
	fmt.Println("CopiedArray:", *copiedArray)

	// now copied array is original array



}



func someFunction()(int,int){
	return 1,2
}
