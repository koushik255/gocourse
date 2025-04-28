package main

import (
	"fmt"
	"time"
)

func main() {

	// Mon Jan 2 15:04:05 MST 2006 /each of these is represented by a specific pattern

	// time parsing patters ISO 8610 format

	layout := "2006-01-02T15:04:05Z07:00"
	str := "2024-07-04T14:30:18Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time!", err)
		return
	}
	fmt.Println(t)

	// consider time zones and that bs when converting into different time zones

	str1 := "Jul 03, 2024 03:18 PM"
	layout1 := "Jan 02, 2006 03:04 PM"

	t1, err := time.Parse(layout1, str1)
	fmt.Println(t1)

}
