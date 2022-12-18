package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Herbert"] = 67

	// switch without argument when not checking a boolean
	switch {
	case ages["Herbert"] < 18:
		fmt.Println("Herbert can't vote.")
	case ages["Herbert"] < 67:
		fmt.Println("Herbert can't retire.")
	case ages["Herbert"] > 67:
		fmt.Println("Herbert can retire.")
	default:
		fmt.Println("Herbert is exactly 67 years old.")
	}

	switch ages["Herbert"] {
	case 1, 2, 3 , 4, 7, 11, 13, 17, 19:
		fmt.Println("Herbert's age is a prime number")
	case 16:
		fmt.Println("Herbert can drive now")
	case 18:
		fmt.Println("Herbert can vote now")
	case 67:
		fmt.Println("Herbert can retire now")
	default:
		fmt.Println("Nothing special about Herbert's age")
	}
}
