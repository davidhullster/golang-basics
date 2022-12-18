package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("We're counting (i):", i)
	}
	a = 0
	for a < 10 {
		fmt.Println("We're counting (a): ", a)
		a++
	}

	a := 0
	for a < 10 {
		if a%2 == 0{
			a++
			continue
	} else if a == 5 {
			break
	}

	fmt.Println("We're counting (a)", a)
	a++
}
