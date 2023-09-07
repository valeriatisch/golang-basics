package main

import (
	"fmt"
)

// Mantra: "Do not communicate by sharing memory, share memory by communicating."

func main() {

	// Channels
	// - A channel is a communication pipe that allows Goroutines to communicate with each other (pass/receive values).
	// - A channel will block by default if it is full.

	// The following won't work
	// ch := make(chan int)
	// ch <- 42 // Channel blocks here, fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(<-ch)

	// Unbuffered channel
	// - An unbuffered channel is a channel that can hold only one value.
	unbufCh := make(chan string)
	// A goroutine needs to be created to send a value to the channel
	// Then the channel blocks / waits.
	go func() {
		unbufCh <- "Willi"
	}()
	// The main-goroutine continues and then receives the value from the channel.
	fmt.Println(<-unbufCh) // Channel blocks here until the value is received from the goroutine above.

	// Buffered channel
	// - A buffered channel is a channel that can hold a certain number of values.
	bufCh := make(chan string, 2)
	bufCh <- "Silvia"
	bufCh <- "Luisa"
	fmt.Println(<-bufCh, ",", <-bufCh)

	bufCh <- "Finn"
	bufCh <- "Luna"
	// The following won't work
	// bufCh <- "Lea" // This is waiting till a value is consumed from the channel so a space becomes free.

	ch := make(chan int)
	go foo(ch, 34)
	go foo(ch, 35)
	fmt.Println(<-ch, ",", <-ch)

	// Ranges
	// - A loop that iterates over a channel until it is closed.
	// - The range will block by default if the channel is empty unless it is closed.
	go func() {
		fmt.Println("Putting values on the channel...")
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	v, ok := <-ch
	fmt.Println(v, ok)

	fmt.Println("Now reading values from the channel...")
	for v := range ch {
		fmt.Println(v)
	}
	v, ok = <-ch
	fmt.Println(v, ok)
}

func foo(ch chan int, x int) {
	ch <- x
}
