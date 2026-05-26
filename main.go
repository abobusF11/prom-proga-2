package main

import "fmt"

func FormatMessage(prefix string, name string) string {
	if name == "" {
		name = "guest"
	}

	return prefix + ", " + name + "!"
}

func Greet(name string) string {
	return FormatMessage("Hello", name)
}

func Farewell(name string) string {
	return FormatMessage("Goodbye", name)
}

func main() {
	fmt.Println(Greet("Git Flow"))
	fmt.Println(Farewell("Git Flow"))
}
