package main

import "fmt"

func main() {
	// var cards string = "Ace of spades"
	// fmt.Println(cards)

	// cardsStr := "Card string" //Declare variable with any type, this call colon-equals
	// fmt.Println(cardsStr)
	// cardsStr = "Card string changed" //Change value of variable
	// fmt.Println(cardsStr)

	// cardsBoolean := false
	// fmt.Println(cardsBoolean)
	// cardsBoolean = true
	// fmt.Println(cardsBoolean)

	// var deckSize int
	// deckSize = 52
	// fmt.Println(deckSize)
	// cards := newCard()
	// fmt.Println(cards)
	// printState()
	//Luu y: printState() thuoc file khac => Muon chay duoc phai dung lenh go run main.go state.go
	// cards := deck{"Ace of spades", newCard(), "Boom"}
	// cards = append(cards, "Spider man")
	// cards.print()
	cards := newDeck()
	cards.print()

	//Deal
	hand, remainingCards := deal(cards, 5)
	fmt.Println("Hand: ")
	hand.print()
	fmt.Println("Remaining cards: ")
	remainingCards.print()

	//Save to file
	// cards.saveToFile("my_cards")
	// cards.newDeckFromFile("my_cards")

	fmt.Println("Before shuffling cards: ", cards)
	cards.shuffleCard(cards)
	fmt.Println("After shuffling cards: ", cards)

	// fruits := []string{"Banana", "Watermelon", "Apple", "Coconut", "Mango"}
	// fmt.Println(fruits[0:2])
	// fmt.Println(fruits[:2])
	// fmt.Println(fruits[2:])

	// greeting := []byte("Hi there") //Type conversion
	// fmt.Println(greeting)

	checkNumber()
}

func newCard() string {
	return "Five of diamonds"
}

func checkNumber() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
