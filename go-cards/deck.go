package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

/* deck is a custom data-type that can be used to create new variables. This variable extends the features of the existing go data-types and can add some custom methods (functions) to it.
This is similar to a class declaration in OOP languages
*/
type deck []string

/* (d deck) is the 'reciever' of the custom method (function) `print()` of our deck custom data-type
1. d references the working copy of the deck variable calling the function print
it is similar to `this` in JS and `self` in python.
By convention in go the name of the ref are either the first letter or first two letters of the custom data-type
deck implies that every variable of type deck can call this function
*/
// This is a reciever function
/*You can use key words as arguements---wiered!!
Functions in go
1. starts with the `func` keyword
2. (ref to variable `d` custom data-type `deck`) recievers in paranthes
3. function name
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*Reciever function to create a new deck of cards*/

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearths", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}

/*You can use key words as arguements---wiered!!
Functions in go
1. starts with the `func` keyword
2. function name
3. (ref to variable `d` custom data-type `deck`, arguments with their data-type) recievers in paranthes
3. function name
4. data-type that specifies the return type which could be any  data-type
if multiple returns put them in parantheses
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",\n")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	stringSlice := strings.Split(string(content), ",\n")
	return deck(stringSlice)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
