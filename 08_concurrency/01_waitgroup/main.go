package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// There is always an init() running before main()
func init() {
	fmt.Println("Initialization...")
}

func main() {
	// Computer specs
	fmt.Println("\nOS 🖥️:\t\t", runtime.GOOS)
	fmt.Println("Arch 🏗️:\t\t", runtime.GOARCH)
	fmt.Println("CPUs 🧠:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines 🚀:\t", runtime.NumGoroutine())

	start := time.Now()

	var wG sync.WaitGroup
	wG.Add(3)

	var userName string
	// Launch a goroutine
	go func() {
		userName = fetchUser()
		fmt.Println("Goroutines:\t", runtime.NumGoroutine())
		wG.Done()
	}()
	// It takes some time to launch a goroutine, during which the main() function continues to run.
	// The main() function might finish before the goroutine does.
	// To avoid this, we can use the WaitGroup.

	var likes int
	go func() {
		likes = fetchUserLikes(userName)
		wG.Done()
	}()

	var match string
	go func() {
		match = fetchUserMatch(userName)
		wG.Done()
	}()

	wG.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			fetchUserLikes(userName)
			wG.Done()
		}()
	}

	wG.Wait()

	fmt.Println("\nUser:", userName)
	fmt.Println("Likes:", likes)
	fmt.Println("Match:", match)

	fmt.Println("\nThat took", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "John"
}

func fetchUserLikes(userName string) int {
	time.Sleep(time.Millisecond * 150) // cannot be faster than 150ms overall
	return 10
}

func fetchUserMatch(userName string) string {
	time.Sleep(time.Millisecond * 100)
	return "Jane"
}
