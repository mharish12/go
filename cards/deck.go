package main

import (
	"fmt"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	suites := []string{"Diamonds", "Spade", "Heart", "Clubs"}

	cardValues := []string{"Ace", "One", "Two", "Three"}

	for _, card := range suites {
		for _, val := range cardValues {
			cards = append(cards, card+" of "+val)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// func deal(d deck, handSize int) (deck, deck) {
// 	return d[:handSize], d[handSize:]
// }

// func (d deck) deal(handSize int) (deck, deck) {
// 	return d[:handSize], d[handSize:]
// }

// func (d deck) toString() string {
// 	return strings.Join([]string(d), "\n")
// }

// func (d deck) saveToFile(fileName string) error {
// 	return os.WriteFile(fileName, []byte(d.toString()), 0666)
// 	// return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
// }

// func readFromFile(fileName string) deck {
// 	bs, er := os.ReadFile(fileName)
// 	if er != nil {
// 		fmt.Println("Error: ", er)
// 		os.Exit(1)
// 	}

// 	fileString := string(bs)

// 	cards := strings.Split(fileString, "\n")

// 	return deck(cards)
// }

func (d deck) shuffle() {

	// n := (int64)(len(d) - 1)
	// source := rand.NewSource(n - 1)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
