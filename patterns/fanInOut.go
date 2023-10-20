package patterns

import (
	"fmt"
	"sync"
)

const (
	maxOrderCount = 3
)

var (
	orderList = []string{"sandwich", "paratha", "frankie"}
)

// Expected Output

// dishRequest -> getPotatoes ->  MakeDish -> Sandwich
// 				  getVeggies  ->		   -> Paratha
//										   -> Frankie

func TryFanInOut() {
	fmt.Printf("\nHi, from Fan-In and Fan-Out\n---------------------------\n\n")
	wg := sync.WaitGroup{}

	dishRequest := make(chan string)
	wg.Add(1)
	go MakeDish(dishRequest, &wg) // spawn the main go-routine

	// now send the orders through dishRequest channel
	for _, order := range orderList {
		wg.Add(1)
		orderName := order
		go func() {
			defer wg.Done()
			dishRequest <- orderName
		}()
	}

	wg.Wait()
}

func MakeDish(dishRequest chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	var dishName string

	myPotatoes := <-getPotato()
	myVeggies := <-getVeggies()

	orderCount := 1

	if myPotatoes && myVeggies { // lets read the values now
		fmt.Printf("Got the potatoes and the veggies!\n")
		for orderCount <= maxOrderCount {
			select {
			case dishName = <-dishRequest:
				fmt.Printf("Order[%s] received!\n", dishName)
				orderCount += 1
				status := processOrder(dishName)
				if <-status {
					fmt.Printf("Order[%s] is ready!\n", dishName)
				}

			default:
				fmt.Printf("Sorry! Can\\'t make %s, only can make Sandwich & Paratha.", dishName)
			}
		}

	}
}

func processOrder(dishName string) <-chan bool {
	fmt.Printf("Order[%s] in process...\n", dishName)
	order := make(chan bool)

	// create its go-routine
	go func() {
		order <- true
	}()

	return order
}

func getPotato() <-chan bool {

	potatoChan := make(chan bool) // receive only channel

	go func() {
		potatoChan <- true
	}()

	// then return the channel
	return potatoChan
}

func getVeggies() <-chan bool {

	veggiesChan := make(chan bool)

	go func() {
		veggiesChan <- true
	}()

	// then return the veggies channel
	return veggiesChan
}
