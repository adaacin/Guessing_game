package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//this is a random number generator
	fmt.Println("This is Ada David-Worlu's Guessing game")
	fmt.Println("You are required to guess the number")

	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(105)
	
	var UserGuess int
	var trials = 7
	var playcount = 1
	for {
		fmt.Println("You can only try", +trials, "time(s)!, so think properly before attempting \n",
			"Am waiting for your input:")
		fmt.Scan(&UserGuess)
		fmt.Println("Note:You have made",
			+playcount, "attempt(s)! ")

		if UserGuess > secretNumber{
			fmt.Println("The number is too high, try again!")	
		} else if UserGuess < secretNumber{
			fmt.Println("The number is too low, try again!")
		} else {
			fmt.Println("And, finally, you guessed correctly, Congratualation!!!")
		}
		playcount++
		trials--
		if playcount > 7 || trials < 1 {
			fmt.Println("Opps sorry, you have exhaused all the trials, start all over!")
			break
		}
	}
}
