package main

import (
	"fmt"
)

type Person struct {
	Name string
}

// Functions

// Methods
// Methods are functions that are attached to a particular type.

func main() {
	// Methods

	// Defer an action
	// Defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.

	// Open and close a file file.txt

	// Variadic parameters allows you to pass an arbitrary number of parameters to a function
	println(sum(1, 3, 4, 5, 6))
	println(sum(3))

	// Anonymous functions
	// An anonymous function is a function without a name.
	func() {
		x := 0
		x += 1
		fmt.Println(x)
	}()

	// With an argument
	func(x int) {
		x += 1
		fmt.Println(x)
	}(0)

	// With a return value
	y := func(x int) int {
		x += 1
		return x
	}(0)
	println(y)

	// Func expressions
	// A func expression is a function without a name that is assigned to a variable.
	// The function can be called using the variable name.
	// The function can be passed as an argument to another function or returned from another function.
	fn := func() {
		fmt.Println("I am a func expression.")
	}
	fn()

	anotherFn := func(x int) string {
		return fmt.Sprintf("%d", x)
	}
	callbackFunc(anotherFn)
}

// Variadic parameters function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Callbacks
// Function taking a function as an argument
func callbackFunc(f func(int) string) {
	f(1)
}

// TODO: Please define a function which will return a func expression.
