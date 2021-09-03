package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a wait group used in schronizing processes
	var wg sync.WaitGroup
	// Create a channel that received boolean when a goroutine is done
	c := make(chan bool)
	// Add the the number of goroutines initialise
	wg.Add(1)
	go func() {
		// called when the goroutine finishes it job
		// or terminated
		defer wg.Done()

		for <-c {
			fmt.Printf("Hello! No Error found \n")

			if !<-c {
				// received a signal that an error has occurred
				fmt.Printf("Oooops! Error found. Shutting down....\n")
			}
		}
	}()
	// simulate a running debugging process where
	// the program hits a bug at 37
	// a ping is sent to the channel to either continue the program
	// or exit
	for i := 0; i < 50; i++ {
		// no error yet
		c <- true
		if (i + 1) == 37 {
			// found an error send signals to the channel
			// to exit an print out a debug message to the console
			close(c)
			break
		}
	}
	// Waits for the goroutine to finish its job
	wg.Wait()
}
