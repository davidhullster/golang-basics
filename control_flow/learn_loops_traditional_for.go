package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("We're counting (i):", i)
	}
	a := 0
	for a < 10 {
		fmt.Println("We're counting (a): ", a)
		a++
	}

	a = 0
	for a < 10 {
		if a%2 == 0{ // if a is even
			a++        // don't print it, but increment by 1
			continue   // go back to top of loop
	} else if a == 5 {  // we only get here if a is not even
			break           // end for loop
	}
  // prints only 1 and 3, since 2, 4, 6, 8 are even
	// 1 is evaluated before the 'break', then the 5, which then runs break
	fmt.Println("We're counting (a)", a)
	a++
}
}
