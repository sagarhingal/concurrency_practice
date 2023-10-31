package cli

import (
	"bufio"
	"concurrency_practice/patterns/fanInOut"
	"concurrency_practice/patterns/rateLimit"
	"concurrency_practice/patterns/workerPool"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

var (
	menuList = []string{
		"Fan-In-Out",
		"Worker-Pool",
		"Rate-Limiting",
		"Circuit-Breaker",
		"Producer-Consumer",
		"Publish-Subscribe",
		"Barrier",
		"Future",
		"Pipeline",
		"Semaphore",
		"For-Select-Done",
		"Try Again OR Clear console",
		"Exit",
	}
)

func Menu() {

	// print the menu
	fmt.Printf("\n------\nMenu -\n------\n")
	for index, menuItem := range menuList {
		fmt.Printf("%d) %s\n", index+1, menuItem)
	}

	// now take the user input
	buf := bufio.NewReader(os.Stdin)
	fmt.Printf("\n> ")
	arg, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		Menu()
	}

	// clean the string
	strArg := strings.Trim(string(arg), "\r\n")
	choice, err := strconv.ParseInt(strArg, 10, 64)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		fmt.Printf("invalid choice: %d\nPlease choose from the following menu:\n\n", choice)
		Menu()
	}

	// now check the choice and call respective menu option
	switch choice {
	case 1:
		// call fanIn-fanOut
		go fanInOut.TrySample()

		// wait for it to finish
		time.Sleep(2 * time.Second)
		Menu()
	case 2:
		// call worker-pool
		go workerPool.TryWorkerPool()

		// wait for it to finish
		time.Sleep(2 * time.Second)
		Menu()
	case 3:
		// call rate-limiting
		rateLimit.TryRateLimit()

		// wait for it to finish
		time.Sleep(1 * time.Second)
		Menu()
	case 4:
		// call circuit-breaker
		fmt.Printf("\nwork in progress......please choose from the menu below:\n")

		// wait for it to finish
		time.Sleep(1 * time.Second)
		Menu()
	case 5:
		// call producer-consumer
		fmt.Printf("\nwork in progress......please choose from the menu below:\n")

		// wait for it to finish
		time.Sleep(1 * time.Second)
		Menu()
	case 6:
		// call publish-subscribe
		fmt.Printf("\nwork in progress......please choose from the menu below:\n")

		// wait for it to finish
		time.Sleep(1 * time.Second)
		Menu()
	case 7:
		// call barrier
		fmt.Printf("\nwork in progress......please choose from the menu below:\n")

		// wait for it to finish
		time.Sleep(1 * time.Second)
		Menu()
	case 8:
		// call future
		fmt.Printf("\nwork in progress......please choose from the menu below:\n")

		// wait for it to finish
		time.Sleep(1 * time.Second)
		Menu()
	case 9:
		// call pipeline
		fmt.Printf("\nwork in progress......please choose from the menu below:\n")

		// wait for it to finish
		time.Sleep(1 * time.Second)
		Menu()
	case 10:
		// call semaphore
		fmt.Printf("\nwork in progress......please choose from the menu below:\n")

		// wait for it to finish
		time.Sleep(1 * time.Second)
		Menu()
	case 11:
		// call for-select-done
		fmt.Printf("\nwork in progress......please choose from the menu below:\n")

		// wait for it to finish
		time.Sleep(1 * time.Second)
		Menu()
	case 12:
		screen.Clear()
		Menu()
	case 13:
		break
	default:
		screen.Clear()
		fmt.Printf("\ninvalid choice: %d\nPlease choose from the following menu:\n", choice)
		Menu()
	}

}
