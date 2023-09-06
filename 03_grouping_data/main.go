package main

import "fmt"

// Structs
// - A struct is a collection of fields.
type Person struct {
	name string
	age  int
}

// Embedded structs
// - A struct can contain other structs.
type Child struct {
	Person
	favToy string
}

func main() {
	// Arrays
	// - An array is a fixed-size collection of elements of the same type.
	// - An array is declared by specifying the type and size of the array.
	// - An array is indexed starting at 0.
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// Length of an array is the number of elements it was declared with.
	fmt.Println(len(arr))

	// Slices
	// - A slice is a dynamically-sized, flexible view into the elements of an array.
	slice := []int{1, 1}
	println(len(slice))
	slice = append(slice, 4, 6, 6)
	slice[1] = 2
	fmt.Println(slice)
	println(len(slice))

	// Composite literal
	// x := type{values}

	// Multi-dimensional slices
	// - A slice can contain any type, including other slices.
	mlt_slice := [][]int{{1, 2, 3}, {4, 5, 6, 5}}
	fmt.Println(mlt_slice)

	// Maps
	// - A map is an unordered collection of key-value pairs.
	// - A map is declared by specifying the type of its key and value.
	// - A map is indexed by its key.
	m := map[string]int{
		"Martin": 42,
		"John":   32,
	}
	fmt.Println(m["Martin"])
	fmt.Println(m["John"])
	// No entry
	fmt.Println(m["Lea"])
	val, ok := m["Lea"]
	fmt.Println(val)
	fmt.Println(ok)

	// Adding new element
	m["Lea"] = 22
	val, ok = m["Lea"]
	fmt.Println(val)
	fmt.Println(ok)

	// Deleting an element
	delete(m, "Lea")

	// Structs
	var p Person = Person{
		name: "Martha",
		age:  62,
	}
	fmt.Println(p.name, p.age)
	p.name = "Martha Stewart"

	// Embedded structs
	var c Child = Child{
		Person: Person{
			name: "Finn",
			age:  4,
		},
		favToy: "Lego",
	}
	fmt.Println(c)

	// Anonymous struct
	anonym := struct {
		name  string
		color string
	}{
		name:  "Apple",
		color: "Green",
	}
	fmt.Println(anonym)
}
