package main

import (
	"fmt"
)

func main() {
	name := "Paul"
	message := getMessage(name)
	fmt.Println(message)
}
func getMessage(name string) string {
	return fmt.Sprintf("Hello %v", name)
}
