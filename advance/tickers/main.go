package main

// produces ticks are regular internvals, tickers are useful for preforming tasks
// on a consitanst schedule, fixed intervals.
// helps maintain a consita t scheduale

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)
	ticker2 := time.NewTicker(2 * time.Second)
	stop2 := time.After(10 * time.Second)

	defer ticker.Stop()
	defer ticker2.Stop()

	// shuts it down because the first ticker is done
	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("tick at :", tick)
		case tick2 := <-ticker2.C:
			fmt.Println("tick2  at :", tick2)
		case <-stop2:
			{
				fmt.Println("Stopping tickers2")
				return
			}
		case <-stop:
			{
				fmt.Println("Stopping ticker")
				return
			}
		}
	}
}

///////++++++++++ SCHEUDLING LOGGING,PERIDOCIC TAKSKS, POLLING FOR UPDATES
// func periodicTask() {
// 	fmt.Println("Prefomring peridodic task at:", time.Now())

// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// func main() {

// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()

// 	for tick := range ticker.C {
// 		fmt.Println("Tick at :", tick)
// 	}

// 	i := 1
// 	for range 5 {
// 		i *= 2
// 		fmt.Println(i)
// 	}

// 	for tick := range ticker.C {
// 		i *= 2
// 		fmt.Println(i)
// 	}

// }
