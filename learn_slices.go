package main

import "fmt"

// slices allow dynamic sizing, remove number from bracket on initialization
// initialize slice with 0 elements with {}
func main() {
	names := []string{}
	names = append(names, "Bob", "Boothe", "Barbara")

	fmt.Println("Names:", names)
	fmt.Println("Length of Names:", len(names))

// pre-allocate 4 elements to the array
// saves work at runtime allocating space on array
// doesn't allow adding more elements than allocated, i.e. 5 will fail
	names2 := make([]string, 4)
	names2[0] = "Mike"
	names2[1] = "Martha"
	names2[2] = "Moose"
	names2[3] = "Megan"
	// this will fail
	//names2[4] = "Megan"

	fmt.Println("Names2:", names2)
}
