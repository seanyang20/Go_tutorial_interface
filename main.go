package main

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func (eb englishBot) getGreeting() string {
	// VERY CUSTOM logic for generating an english greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
