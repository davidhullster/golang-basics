package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Carl"] = 11
	ages["Cindy"] = 28
	ages["Charles"] = 67
	ages["Candice"] = 18
	ages["Ciara"] = 16

	for name, age := range ages {
		switch age {
		case 1, 2, 3, 5, 7, 11, 13, 17, 19:
			fmt.Println(fmt.Sprintf("%s's age is a prime number.", name))
		case 16:
			fmt.Println(name, "can drive now.")
		case 18:
			fmt.Println(name, "can vote now.")
		case 67:
			fmt.Println(name, "can retire now.")
		default:
			fmt.Println(name, "has nothing special about their age.")
		}
	}
}
