package main

import "fmt"

func main() {
	// For loop: the only one with semicolons
	for i := 0; i < 5; i++ {
		fmt.Println("Counter:", i)
	}

	// TODO: Create slice and for loop through it
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice); i++ {
		fmt.Println("Slice value", slice[i], "at position", i)
	}

	// Range
	for i, val := range slice {
		fmt.Println("Slice value ", val, "at position", i)
	}

	// "While" loop
	i := 0
	for i < 5 {
		fmt.Println("Counter:", i)
		i++
	}

	// Infinite loops
	// for ;; {
	// }

	// for true {
	// }

	// for {
	// }

	// If-else
	if 3 < 4 {
		fmt.Println("3 is less than 4")
	} else {
		fmt.Println("3 is not less than 4")
	}

	// Two statements in one line
	if z := 4; (z < 5 && z > 0) || true {
		fmt.Println(z)
		fmt.Println("z is less than 5")
		z = z + 3
	} else {
		fmt.Println("z is not less than 5")
	}

	// Switch
	var color string = "red"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	}

	// Switch with fallthrough
	// - fallthrough will execute the next case even if it doesn't match
	color = "dark green"
	switch color {
	case "dark green":
		fmt.Println("Color is dark green")
		fallthrough
	case "green":
		fmt.Println("Color is also just green")
	}
}
