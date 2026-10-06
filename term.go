package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	fmt.Printf("Width (columns): %d\n", width)
	fmt.Printf("Height (rows): %d\n", height)
}
