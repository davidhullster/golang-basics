package main

import (
	"bufio"
	"fmt"
	"motd/message"
	"os"
	"strings"
)
// functions must start with a capital letter to be accessible publically
func main() {
	f, err := os.OpenFile("/etc/motd", os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error: Unable to open /etc/motd")
		os.Exit(1)
	}

	defer f.Close() // closes after context

  reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	message := message.Greeting(name, phrase)
	fmt.Println(message)

	_, err = f.Write([]byte(message))

	// err := ioutil.WriteFile("/etc/motd", []byte(message), 0644)

	if err != nil {
		fmt.Println("Error: Failed to write /etc/motd.")
		os.Exit(1)
	}
}
