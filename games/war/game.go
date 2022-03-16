package war

import "cards-server/card"

type player struct {
	username   string
	deck       *card.Deck
	activeCard card.Card
}

type container struct {
	players []player
}

func New(players []string) *container {
	d := card.StandardDeck().Shuffle()
	cardsPerHand := d.Len() / len(players)
	hands := d.Deal(cardsPerHand, len(players))

	c := &container{
		players: make([]player, len(players)),
	}

	for i := 0; i < len(hands); i++ {
		c.players[i] = player{
			deck:     &card.Deck{Cards: hands[i], Discard: make([]card.Card, 0)},
			username: players[i],
		}
	}

	return c
}

func compareCards(players []player) []player {
	topCard := card.Card{Val: card.Joker}
	winning := []int{}

	for i, p := range players {
		if p.activeCard.Val >= topCard {
			winning = append(winning, i)
		} else {

		}
	}
}
