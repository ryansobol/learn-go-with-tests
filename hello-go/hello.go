package main

import "fmt"

func Greet(name string, language string) string {
	var person string

	if name == "" {
		person = "world"
	} else {
		person = name
	}

	return greetingPrefix(language) + ", " + person + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = "Hola"
	case "French":
		prefix = "Bonjour"
	default:
		prefix = "Hello"
	}
	return
}

func main() {
	fmt.Println(Greet("Jesten", ""))
}
