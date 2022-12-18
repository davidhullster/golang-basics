package main

import "fmt"

func main() {
	var myInt int = 16
	var val, ok = "yes", true
	var _ = "not used"

	fmt.Println("myInt is:", myInt)
	fmt.Println("myInt times two:", myInt*2)
	fmt.Println("myInt times two:", myInt*2)
	fmt.Println("val is:", val)
	fmt.Println("ok is:", ok)
}
