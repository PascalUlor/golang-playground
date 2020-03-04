package main

import "fmt"

// custom interface type
type reusable interface {
	greeting() string
}

// create custom struct types
type englishBot struct{}
type spanishBot struct{}

func main()  {
	en := englishBot{}
	sp := spanishBot{}

	prnitGreet(en)
	prnitGreet(sp)
}

// reusable interface function
func prnitGreet(r reusable)  {
	fmt.Println(r.greeting())
}

func (englishBot) greeting() string {
	return "Hello Lady"
}

func (spanishBot) greeting() string {
	return "Holla Seniorita"
}