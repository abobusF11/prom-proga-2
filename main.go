package main

import "fmt"

func Greet(name string) string {
	if name == "" {
		return "Hello, guest!"
	}
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Greet("Git Flow"))
}
