package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("'detect' or 'to' subcommands are required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "detect":
		//call detect.go
	case "to":
		//call to.go
	default:
		fmt.Println("'detect' or 'to' subcommands are required")
		os.Exit(1)
	}
}
