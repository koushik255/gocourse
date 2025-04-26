package main
import (
	"fmt"
)

// stucrst in go
type Person struct {
	firstName string
	lastName string
	age int
	address Address // embedding struct into this structs
	PhomeHomeCell // embedding anonymous field
}

type Address struct {
	city string
	country string
}

type PhomeHomeCell struct {
	home string
	cell string
}

///methods are not just associated with a struct
// they can be associated with any specific type


// declaring methods outside of the struct
// the fullname function is associated with (p Person) the struct

// the a variable is not realated to the one declared in the main function
// its just a parameter just think of a normal function its basically just that but
// only with thign of the person struct cause remeber you hace to call the method on something
// from the struct type so for example you have to do p.fullname and p is going to be of the struct
//Person
func (a Person) fullName() string {
	return a.firstName + " " + a.lastName
}

// no return type since we are modifying a value
// make the memory address availble to the function
// if we just pass a value it will not modify the original variable

func (a *Person) incrementAgeByOne() {
	a.age++
}

// go allows for mebeeding structs within other structs


func main() {
	
	
	p := Person {
		firstName: "Koushik",
		lastName: "Kheda",
		age: 23,
		address: Address{
			city: "London",
			country: "England",
		},
		PhomeHomeCell: PhomeHomeCell{
		home: "5199000000",
		cell: "2265192450",
		},
	}
	// instances of structs
	p1:= Person {
		firstName: "Prabir",
		age: 233,
	}

	p1.address.city = "New york"
	p1.address.country = "USA"
	
	fmt.Println(p.firstName,p.lastName,p.age)
	fmt.Println(p1.firstName,p1.lastName,p1.age)

	// Anonymous Structs 
	// user := struct {
	// 	username string
	// 	email string
	// }{
	// 	username : "user123",
	// 	email: "pseudoemail@example.org",
	// }

	fmt.Println("********NEW*********")
	// methods in structs
	// generated the full name of p
	fmt.Println(p.fullName())
	fmt.Println(p.address.city)
	fmt.Println(p.address.country)
	fmt.Println(p.PhomeHomeCell.home)
	fmt.Println(p.cell)




	fmt.Println(p1.fullName())
	fmt.Println(p.address.city)
	fmt.Println(p.address.country)

	fmt.Println(p.age)
	p.incrementAgeByOne()
	fmt.Println(p.age)


}



