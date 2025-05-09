package main

import (
	"fmt"
	"sort"
)

// sorting is the process of arranging data in a specific order

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool //
}

func (s *personSorter) Len() int {
	return len(s.people)
}

func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

// this code is just reapeating the code below

func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
	// since the personSorter implements the s3 things for the sort interface less swap len it can
	// be passed as a interface to teh sort.Sort function
}

// type ByAge []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// 	// we are comparing ages
// 	// this is basicaly just sorting thoruhg the ages
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// 	// this is swapping the values
// }

func main() {
	people := []Person{
		{"John", 30},
		{"Jane", 25},
		{"Jim", 35},
		{"Jill", 28},
	}
	fmt.Println("UNSorted people: ", people)
	// 	sort.Sort(ByAge(people))
	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	By(age).Sort(people)
	fmt.Println("Sorted people: ", people)

	// numbers := []int{5, 3, 4, 2, 1}
	// sort.Ints(numbers) // sorts the numbers in ascending order
	// // it modifys the original slice
	// fmt.Println("Sorted numbers: ", numbers)

	// stringSlice := []string{"banana", "apple", "cherry"}
	// sort.Strings(stringSlice) // sorts the strings in ascending order
	// fmt.Println("Sorted strings: ", stringSlice)

}
