package patterns

import (
	"fmt"
	"time"
)

func TryFanInOut() {
	fmt.Printf("\nHi, from Fan-In and Fan-Out\n")

	dishRequest := make(chan string)

	MakeDish(&dishRequest)

	go func() {
		dishRequest <- "sandwich"
	}()

	go func() {
		dishRequest <- "paratha"
	}()
}

func MakeDish(dishRequest *chan string) {

	myPotatoes := <-getPotato()
	myVeggies := <-getVeggies()
	dishName := <-*dishRequest

	if myPotatoes && myVeggies { // lets read the values now
		switch dishName {
		case "sandwich":
			makeSandwich()
		case "paratha":
			makeParatha()
		default:
			fmt.Printf("Sorry! Can\\'t make %s, only can make Sandwich & Paratha.", dishName)
		}
	}
}

func makeSandwich() chan bool {
	sandwich := make(chan bool)

	sandwich <- true

	return sandwich
}

func makeParatha() chan bool {
	paratha := make(chan bool)

	paratha <- true

	return paratha
}

func getPotato() <-chan bool {

	potatoChan := make(chan bool) // receive only channel

	go func() {
		potatoChan <- true

		// lets give it some time
		time.Sleep(100 * time.Millisecond)
	}()

	// then return the channel
	return potatoChan
}

func getVeggies() <-chan bool {

	veggiesChan := make(chan bool)

	go func() {
		veggiesChan <- true

		// lets give this one some time as well
		time.Sleep(100 * time.Millisecond)
	}()

	// then return the veggies channel
	return veggiesChan
}
