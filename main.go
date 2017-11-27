package main

import "fmt"

type bot interface{
	getGreeting() string // all structs that contain getGreeting and return a string become members of type bot
}
type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string{
	// VERY custom logic for generating an english greeting
	return "Hello"
}

func (spanishBot) getGreeting() string{
	// VERY custom logic for generating a spanish greeting
	return "Hola"
}

