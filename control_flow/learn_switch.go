package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Herbert"] = 67

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
}
