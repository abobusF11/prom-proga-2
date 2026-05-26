package main

import "fmt"

// FormatMessage returns a formatted message with the given prefix and name.
func FormatMessage(prefix string, name string) string {
	if name == "" {
		name = "guest"
	}

	return prefix + ", " + name + "!"
}

// Greet returns a greeting message for the given name.
func Greet(name string) string {
	return FormatMessage("Hello", name)
}

// Farewell returns a goodbye message for the given name.
func Farewell(name string) string {
	return FormatMessage("Goodbye", name)
}

func main() {
	fmt.Println("Message from developer two")
}
