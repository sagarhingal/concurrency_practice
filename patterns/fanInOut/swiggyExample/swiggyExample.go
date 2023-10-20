package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	item string
	id   int
	src  string
}

func newOrder(items ...Order) <-chan Order {
	out := make(chan Order)
	go func() {
		for _, n := range items {
			out <- n
		}
		close(out)
	}()
	return out
}

func main() {
	//restaurent with n cooks
	TableOrder := newOrder(Order{item: "Pizza", id: 1, src: "waiter1"},
		Order{item: "Paratha", id: 2, src: "waiter1"},
		Order{item: "Tea", id: 3, src: "waiter1"},
		Order{item: "Tea", id: 3, src: "waiter1"},
		Order{item: "Tea", id: 3, src: "waiter1"},
		Order{item: "Tea", id: 3, src: "waiter1"},
		Order{item: "Pizza", id: 4, src: "waiter1"},
		Order{item: "Tea", id: 5, src: "waiter1"},
	)

	DriveINOrder := newOrder(Order{item: "Paratha", id: 6, src: "CAR1546"},
		Order{item: "Paratha", id: 7, src: "CAR1546"},
		Order{item: "Pizza", id: 8, src: "CAR1546"},
		Order{item: "Pizza", id: 9, src: "CAR1546"},
		Order{item: "Paratha", id: 10, src: "CAR1546"},
	)

	OnlineOrder := newOrder(Order{item: "Pizza", id: 11, src: "SWIGGY1546"},
		Order{item: "Pizza", id: 12, src: "SWIGGY1546"},
		Order{item: "Pizza", id: 13, src: "SWIGGY546"},
		Order{item: "Paratha", id: 14, src: "SWIGGY46"},
		Order{item: "Paratha", id: 15, src: "SWIGGY546"},
		Order{item: "Pizza", id: 16, src: "SWIGGY1546"},
		Order{item: "Pizza", id: 17, src: "SWIGGY546"},
		Order{item: "Paratha", id: 18, src: "SWIGGY46"},
		Order{item: "Paratha", id: 19, src: "SWIGGY546"},
	)

	// Consume the merged output from c1 and c2.
	for n := range ProcessOrder(TableOrder, DriveINOrder, OnlineOrder) {
		fmt.Println(n)
	}
}

func ProcessOrder(cs ...<-chan Order) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	output := func(c <-chan Order) {
		for n := range c {
			switch n.item {
			case "Tea":
				time.Sleep(1 * time.Second)
			case "Pizza":
				time.Sleep(2 * time.Second)
			case "Paratha":
				time.Sleep(3 * time.Second)
			}
			out <- n.item + " ready for " + n.src
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
