package main

import (
	"fmt"
	"math/rand"
)

// seed is a starting point in generating a random number
// seeds are used to init the sequence
// seeds are good for reproducabillty, same as minecraft seed blud

func main() {
	fmt.Println("Hello my name is the shik")
	fmt.Println(rand.Intn(100) + 1) // random numbers from 0 -101 seeds automatically
	// half open interval , [0,n) half open interval just think domain and range like sqaure bracket is not includeing in the thing

	// number := 0
	// for number != 30 {
	// 	number = rand.Intn(200)
	// 	fmt.Println("Generated number:", number)

	// }
	// fmt.Println("Generated 101! loop is done!", number)

	number1 := 0
	for number1 != 30 {
		number1 = rand.Intn(101)
		fmt.Println("Generated number:", number1)
	}
	fmt.Println("GENERATED NUMBER 30", number1)

	for {
		// show the menu
		fmt.Println("Welcome to the dice game")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. exit")
		fmt.Print("Enter your choice!")

		var choice int
		_, err := fmt.Scan(&choice) // pass the memory addres of chioce
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("INVALID CHOICE PLEASE ENTER 1 or 2")
			continue
		}
		if choice == 2 {
			fmt.Println("thank you for playing the game!")
			break
		}

		// Show reults
		// fmt.Printf("You rolled a %d and a %d.\n", die1, die2)
		// fmt.Println("The total is:", total)
		total := 0
		attemps := 0
		fmt.Println("Rolling will stop once you roll a 12!")
		for total != 12 {
			die1 := rand.Intn(6) + 1
			die2 := rand.Intn(6) + 1
			total = die2 + die1
			fmt.Printf("Die 1 : %d, Die 2 : %d\n", die1, die2)
			fmt.Println("Total is :", total)
			attemps++
			fmt.Println(attemps)

		}
		fmt.Println("You got a perfect 12!", total)

		// ask if the user wants to roll the dice again
		fmt.Print("Do you want to roll again?(y/n)")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("INVALID INPUT INPUT y or n Assuming no.")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thank you for playing the game!")
			break
		}
	}

}
