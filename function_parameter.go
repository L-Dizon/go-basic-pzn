package main

import "fmt"

func main() {
	sayHelloTo("say", "gedomazo")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
