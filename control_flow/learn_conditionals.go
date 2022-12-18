package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Larry"] = 1

	if ages["Larry"] < 18 {
		fmt.Println("Larry can't vote.")
	} else if ages["Larry"] < 67{
		fmt.Println("Larry is not retirement age.")
	} else {
		fmt.Println("Larry can retire.")
	}
}
