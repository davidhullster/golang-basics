package main

import "fmt"

func main() {
	birthdays := map[string]string {
		"Garry": "09/21/1999",
		"Greg": "08/30/2012",
		"Gloria": "01/14/2005",
	}
	delete(birthdays, "Greg")
	fmt.Println("Birthdays:", birthdays)

	ages := map[string]int{}
	ages["Garry"] = 28
	ages["Greg"] = 23
	ages["Gloria"] = 57

	fmt.Println("Ages:", ages)
}
