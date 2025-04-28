package main

import (
	"fmt"
	"regexp"
)

// for when things need to be matched or whatever

func main() {
	fmt.Println(`He said, "I am great"`)

	//compile a regex of the format which is of a email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`) // + means anything in bracket cna happen multiple times
	// hyphen must be on the end because its used as a range characters
	//testing strings
	email := "user@email.com"
	email2 := "Invalid_email"

	//match
	fmt.Println("Email 1 :", re.MatchString(email))
	fmt.Println("Email 2 :", re.MatchString(email2))

	if re.MatchString(email2) {
		return
	} else {
		fmt.Println("ERROR BLUD ON EMAIL2")
	}

	// Capturing Groups with regexp
	// compile a regex pattern to capture date components
	re = regexp.MustCompile(`(\d){4}-(\d{2})-(\d{2})`)

	// test string
	date := "2024-07-30"

	// find all submatches

	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

}
