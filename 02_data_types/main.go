package main

import "fmt"

func main() {
	// --- Variables ---
	// Go is a statically typed language, but it is able to figure out the type for you

	// Declare a variable & assign a value to it with :=
	// Btw: The variables we declare must always be used
	a := 3
	println(a)

	// Re-assign the variable
	a = 4

	// Can't assign a value to a variable with a different type
	// x := "hello"
	// x = false

	// Declare a variable and assign a value with var
	var b = 5
	println(b)
	fmt.Printf("%T \n", b)

	var c int = 10
	println(c)

	// Declare multiple variables with a type at once
	var (
		d,
		e,
		f int
	)
	println(d, e, f)

	// Shadowing
	globalVar = 1
	println(globalVar)

	// Zero value: default value for a variable
	// int: 0
	// float: 0.0
	// string: ""
	// bool: false
	// nil: functions, pointers, slices etc.
	t := true
	println(t)

	var fa bool
	println(fa)

	// A string is immutable, you can assign a new value but can't change the bytes
	s := "hello"
	println(s)
	s = "world"
	println(s)

	// --- Custom Types ---
	// Creating your own type
	type myOwnInteger int

	var g myOwnInteger = 10
	println(g)
	fmt.Printf("%T \n", g)

	// Conversion
	h := int(g)
	println(h)
	fmt.Printf("%T \n", h)

	// --- Logical Operators ---
	// &&, ||, !
	var fls = false
	var tru = true
	println(fls && tru)
	println(fls || tru)
	println(!tru)

	// --- Constants ---
	// Typed
	const i int = 10

	// Define multiple constants
	// Untyped
	const (
		j = 10
		k = 20
	)

	// Iota
	type Weekday int
	const (
		monday Weekday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

}

// Can't do this in outer scope/package-level:
// x := 1
// but this is ok:
var globalVar = 2

func f() {
	println(globalVar)
}
