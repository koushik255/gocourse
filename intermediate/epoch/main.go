package main

import (
	"fmt"
	"time"
)

// refers to a point in time that serves as a timestamp for computing
// its a defined starting point
// epoc is unix epoc with is 00:00:00 UTC on January 1st 1970

func main() {
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time:", unixTime)
	t := time.Unix(unixTime, 0)
	fmt.Println(t)

}
