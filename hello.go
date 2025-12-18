package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(10) + 1 // random number 1-10
	var guess int

	fmt.Println("Guess the number between 1 and 10!")
	for {
		fmt.Print("Your guess: ")
		fmt.Scanln(&guess)

		if guess == secret {
			fmt.Println("🎉 You got it!")
			break
		} else if guess < secret {
			fmt.Println("Too low, try again!")
		} else {
			fmt.Println("Too high, try again!")
		}
	}
}
