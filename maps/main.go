package main
import (
 "fmt"
"maps"
)

//like dictionaires in python
//provide a good way for key value pairs
//keys are from a compareable type, and is unorderd 



func main(){

	// var m map [key]valueType

	// mapVariale = make(map[keyType]valueType)

	//init using map literal
	// mapVariable = map[keyType]valueType{
	// 	key1 :value1,
	// 	key2 :value2,

	// }

	myMap := make(map[string]int) ///string keys value integers
	fmt.Println(myMap)

	myMap["key1"] = 9 
	myMap ["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	myMap["code"] = 21
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap ["key2"] = 12
	myMap ["key3"] = 15
	myMap ["key4"] = 111
	fmt.Println(myMap)

	//remove all value pairs
	// clear(myMap)
	// fmt.Println(myMap)

	_,unknownvalue := myMap["key2"]
	// fmt.Println(value)
	fmt.Println("is value assoicated with this key pair?:", unknownvalue) /// returns true because there is an key associated with the "key2"
	// the 2nd value checks if there is any value associated with the key
	//you dont have to use this value but for trouble shooting its good

	//int map in a different way // Declare and init on the same line 
	myMap2 := map[string]int {"a": 1, "b": 2}
	fmt.Println(myMap2)

	//Equalityt checks
	
	myMap3 := map[string]int {"a": 1, "b": 2, "C": 3}


	if maps.Equal(myMap3,myMap2){
		fmt.Println("The maps are equal toeachother!")
	} else {
		fmt.Println("UN EQUAL MAPS ABORT MISSON!!")
	}

	// iterate over a map
	for k,v := range myMap3{
		fmt.Println(k,v)
	}

	var myMap4 map[string]string
	myMap4 = make(map[string]string)
	myMap4["Key1"] = "Hello"
	fmt.Println(myMap4)
	delete(myMap4,"Key1")
	fmt.Println(myMap4)
	myMap4["NAME:"]= "Koushik!"
	fmt.Println(myMap4)


}