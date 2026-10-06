package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func main() {
	// 1. Create a channel to receive OS signals
	sigChan := make(chan os.Signal, 1)

	// 2. Register the channel to listen for SIGWINCH (Window Changed)
	signal.Notify(sigChan, syscall.SIGWINCH)

	fmt.Println("Listening for terminal resize events. Resize your window... (Ctrl+C to exit)")

	// 3. Block and listen for events in a loop
	for sig := range sigChan {
		switch sig {
		case syscall.SIGWINCH:
			// Get the updated dimensions
			width, height, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				fmt.Println("Error getting terminal size:", err)
				continue
			}
			fmt.Printf("Terminal resized! Width: %d, Height: %d\n", width, height)

		case syscall.SIGINT:
			fmt.Println("\n Caught Ctrl+C (SIGINT). Cleaning up and exiting...")
			os.Exit(0)
		case syscall.SIGTSTP:
			fmt.Println("\n⏸️Cught Ctrl+Z (SIGTSTP). Custom pause logic goes here.")
			fmt.Println("Note: Process will not enter the background unless you explicitly allow it.")

		default:
			fmt.Printf("unkown signal!\n")
		}
	}
}
