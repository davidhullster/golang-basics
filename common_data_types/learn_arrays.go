package main

import "fmt"

func main() {
	names := [3]string{"Bob", "Barbara", "Boothe"}

	fmt.Println("Names:", names)


	var names2 [3]string
	names2[0] = "Bob"
	names2[1] = "Barbara"
//	names2[2] = "Boothe"

	fmt.Println("Names2:", names2)

	fmt.Println("names2[2] is nil:", names2[2] == "")
}
