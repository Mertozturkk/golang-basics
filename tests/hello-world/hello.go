package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

/*
	func Hello(name string, language string) string {
		if name == "" {
			name = "World"
		}
		if language == spanish {
			return spanishHelloPrefix + name
		}
		if language == french {
			return frenchHelloPrefix + name
		}
		return englishHelloPrefix + name
	}
*/
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
func main() {
	fmt.Println(Hello("Mert", "French"))
}
