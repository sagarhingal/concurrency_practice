package patterns

import (
	"fmt"
	"sync"
	"time"
)

func TryFanInOut() {
	fmt.Printf("\nHi, from Fan-In and Fan-Out\n")
	wg := sync.WaitGroup{}

	dishRequest := make(chan string)

	wg.Add(1)
	go MakeDish(&dishRequest, &wg)

	dishRequest <- "sandwich"

	dishRequest <- "paratha"

	wg.Wait()
}

func MakeDish(dishRequest *chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	dishName := <-*dishRequest
	fmt.Printf("Order received for %s!\n", dishName)
	myPotatoes := <-getPotato()
	myVeggies := <-getVeggies()

	if myPotatoes && myVeggies { // lets read the values now
		fmt.Printf("Got the potatoes and the veggies!\n")
		switch dishName {
		case "sandwich":
			status := <-makeSandwich()
			if status {
				fmt.Printf("Order[%s] is ready!\n", dishName)
			}
		case "paratha":
			status := <-makeParatha()
			if status {
				fmt.Printf("Order[%s] is ready!\n", dishName)
			}
		default:
			fmt.Printf("Sorry! Can\\'t make %s, only can make Sandwich & Paratha.", dishName)
		}
	}
}

func makeSandwich() <-chan bool {
	fmt.Printf("Order[Sandwich] in process...\n")
	sandwich := make(chan bool)

	sandwich <- true

	return sandwich
}

func makeParatha() <-chan bool {
	fmt.Printf("Order[Paratha] in process...\n")
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
