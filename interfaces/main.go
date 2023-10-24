package main

import "fmt"

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (e englishBot) getGreeting() string {
	return "Hello There!"
}

func (e spanishBot) getGreeting() string {
	return "Hola!"
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}
