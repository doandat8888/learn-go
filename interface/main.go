package main

import "fmt"

//Định nghĩa type chung cho toàn ứng dụng
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	ebBot := englishBot{}
	spBot := spanishBot{}
	printGreeting(ebBot) //epBot sẽ có type là bot
	printGreeting(spBot)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
