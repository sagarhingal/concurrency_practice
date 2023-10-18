package main

import (
	"concurrency_practice/cli"
	"fmt"
	"os"
)

func main() {

	// Print the menu
	cli.Menu()
	fmt.Printf("\nGoodbye!\n\n")
	os.Exit(1)
}
