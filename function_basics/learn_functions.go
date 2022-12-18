package main

import "fmt"

func main() {
	message := greeting("Mary", "Hello")
	fmt.Println(message)
	message2 := same_as_greeting("Mary", "Hello")
	fmt.Println(message2)
}

func greeting(name string, message string) string{
	return fmt.Sprintf("%s, %s", message, name)
}

func same_as_greeting(name, message string) (salutation string) {
	salutation = fmt.Sprintf("%s, %s", message, name)
	return
}
