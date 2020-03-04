package main

import "fmt"

func main() {
	var cards string = "God is Great"
	num := 23
	boolean := false
	num = 67
	sampleSlice := []string{"ace"}
	/* create new variables dataStore and dataSlice of type deck which is the custome variable I created. This is similar to creating an instance of a class in OOP languages */
	// dataStore := deck{"hearts", "joker"} // first call of custome data-type deck
	// dataSlice := deck{"ace"} // second call of custome data-type deck
	// dataSlice = append(dataSlice, "jack")
	dataStore := newDeck()
	dataSlice := newDeckFromFile("sample_cards")

	// for i, card := range (dataSlice) {
	// 	fmt.Println(i, card)
	// }
	// each declaration of the deck data-type is exclusive can call functions of the deck data type
	// dataStore.print()
	dataSlice.print()
	// hand, remainingCards := deal(dataStore, 5)
	// hand.print()
	dataStore.saveToFile("sample_cards")
	fmt.Println(newCard(cards), "\n", "Number===",num, boolean, newCard("Diamond"), sampleSlice)
	// remainingCards.print()
	dataSlice.shuffle()
	dataSlice.print()
}

/*You can use key words as arguements---wiered!!
Functions in go
1. starts with the `func` keyword
2. newCard is function name
3. item is function argument
4. data-type that specifies the return type which could be any  data-type
if multiple returns put them in parantheses
*/
func newCard (item string) string {
	return item
}
