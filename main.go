package main

import "fmt"

func Greet(name string) string {
	if name == "" {
		return "Hello, guest!"
	}
	return "Hello, " + name + "!"
}

func Farewell(name string) string {
	if name == "" {
		return "Goodbye, guest!"
	}

	return "Goodbye, " + name + "!"
}

func main() {
	fmt.Println(Greet("Git Flow"))
}
