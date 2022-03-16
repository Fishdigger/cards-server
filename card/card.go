package card

import "fmt"

type Suit string

const (
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Spades   Suit = "Spades"
	Clubs    Suit = "Clubs"
)

type FaceValue int

const (
	Joker FaceValue = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit Suit
	Val  FaceValue
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %v", c.Val, c.Suit)
}

func (v FaceValue) String() string {
	switch v {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Joker:
		return "Joker"
	default:
		panic("unknown face value")
	}
}
