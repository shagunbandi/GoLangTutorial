package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type hindiBot struct{}

func main() {
	eb := englishBot{}
	hb := hindiBot{}

	printGreeting(eb)
	printGreeting(hb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (hindiBot) getGreeting() string {
	return "Namaste"
}
