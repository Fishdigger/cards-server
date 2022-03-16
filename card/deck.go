package card

import (
	"math/rand"
	"strings"
	"time"
)

type Deck struct {
	Cards   []Card
	Discard []Card
}

func (d *Deck) String() string {
	cards := make([]string, len(d.Cards))
	for i, card := range d.Cards {
		cards[i] = card.String()
	}
	return strings.Join(cards, "\n")
}

func (d *Deck) Len() int {
	return len(d.Cards)
}

func StandardDeck() *Deck {
	cards := make([]Card, 0)
	for i := 1; i <= 13; i++ {
		v := FaceValue(i)
		cards = append(cards, []Card{
			{Suit: Spades, Val: v},
			{Suit: Hearts, Val: v},
			{Suit: Clubs, Val: v},
			{Suit: Diamonds, Val: v},
		}...)
	}
	return &Deck{Cards: cards}
}

func (d *Deck) Shuffle() *Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	return d
}

func (d *Deck) Draw() Card {
	top := d.Cards[0]
	d.Cards = d.Cards[1:]
	return top
}

func (d *Deck) Deal(cardsPerHand, numHands int) [][]Card {
	hands := make([][]Card, numHands)
	for j := 0; j < cardsPerHand; j++ {
		for i := 0; i < numHands; i++ {
			hands[i] = append(hands[i], d.Draw())
		}
	}

	return hands
}
