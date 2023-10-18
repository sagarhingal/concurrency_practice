package main

import (
	"bufio"
	"concurrency_practice/patterns"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Print the menu
	menu()
	fmt.Printf("\nGoodbye!\n\n")
	os.Exit(1)
}

func menu() {

	// print the menu
	fmt.Printf("\n------\nMenu -\n------\n1) Fan-In & Fan-Out\n2) Worker-Pool\n3) Try Again\n4) Exit\n")
	// now take the user input
	buf := bufio.NewReader(os.Stdin)
	fmt.Printf("\n> ")
	arg, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		menu()
	}

	// clean the string
	strArg := strings.Trim(string(arg), "\r\n")
	choice, err := strconv.ParseInt(strArg, 10, 64)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		fmt.Printf("invalid choice: %d\nPlease choose from the following menu:\n\n", choice)
		menu()
	}

	// now check the choice and call respective menu option
	switch choice {
	case 1:
		// call fanIn-fanOut
		go patterns.TestPatterns()
		time.Sleep(500 * time.Millisecond)
		menu()
	case 2:
		// TODO: call worker-pool
		fmt.Printf("\nWorker-pool(2) pattern is yet to be implemented...\nPlease choose from these options:\n")
		menu()
	case 3:
		menu()
	case 4:
		break
	default:
		fmt.Printf("invalid choice: %d\nPlease choose from the following menu:\n", choice)
		menu()
	}

}
