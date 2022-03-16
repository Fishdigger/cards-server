package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeck(t *testing.T) {
	t.Run("StandardDeck", func(t *testing.T) {
		deck := StandardDeck()
		assert.Equal(t, 52, deck.Len())
	})

	t.Run("Draw", func(t *testing.T) {
		deck := StandardDeck()
		ogLen := deck.Len()
		deck.Shuffle()

		_ = deck.Draw()

		assert.Equal(t, ogLen-1, deck.Len())
	})

	t.Run("Deal", func(t *testing.T) {
		deck := StandardDeck().Shuffle()
		hands := deck.Deal(2, 2)

		assert.Equal(t, len(hands), 2)
		for _, hand := range hands {
			assert.Equal(t, 2, len(hand))
		}
	})
}
