package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddCard(t *testing.T) {
	expectedCardName := `testCard`
	expectedCard := NewCard(&expectedCardName, 10, `testSuit`)

	deck := NewDeck(nil)
	cards := deck.GetCards()
	require.Len(t, cards, 0)

	deck.AddCard(expectedCard)
	cards = deck.GetCards()
	require.Len(t, cards, 1)
	assert.Equal(t, expectedCard, cards[0])
}

func TestCount(t *testing.T) {
	expectedCardName := `testCard`
	expectedCard := NewCard(&expectedCardName, 10, `testSuit`)

	deck := NewDeck(nil)
	assert.Equal(t, deck.Count(), 0)

	deck.AddCard(expectedCard)
	assert.Equal(t, deck.Count(), 1)
}

func TestDrawBottomN(t *testing.T) {
	cards := []Card{
		NewCard(nil, 1, ``),
		NewCard(nil, 2, ``),
		NewCard(nil, 3, ``),
	}

	deck := NewDeck(nil)
	for _, c := range cards {
		deck.AddCard(c)
	}

	actualCards := deck.DrawBottomN(0)
	require.Len(t, actualCards, 0)
	assert.Equal(t, 3, deck.Count())

	deck = NewDeck(nil)
	for _, c := range cards {
		deck.AddCard(c)
	}

	actualCards = deck.DrawBottomN(1)
	require.Len(t, actualCards, 1)
	actualCard := actualCards[0]
	assert.Equal(t, 3, actualCard.GetValue())
	assert.Equal(t, 2, deck.Count())

	deck = NewDeck(nil)
	for _, c := range cards {
		deck.AddCard(c)
	}

	actualCards = deck.DrawBottomN(2)
	require.Len(t, actualCards, 2)
	actualCard = actualCards[0]
	assert.Equal(t, 2, actualCard.GetValue())
	actualCard = actualCards[1]
	assert.Equal(t, 3, actualCard.GetValue())
	assert.Equal(t, 1, deck.Count())

	deck = NewDeck(nil)
	for _, c := range cards {
		deck.AddCard(c)
	}

	actualCards = deck.DrawBottomN(3)
	require.Len(t, actualCards, 3)
	actualCard = actualCards[0]
	assert.Equal(t, 1, actualCard.GetValue())
	actualCard = actualCards[1]
	assert.Equal(t, 2, actualCard.GetValue())
	actualCard = actualCards[2]
	assert.Equal(t, 3, actualCard.GetValue())
	assert.Equal(t, 0, deck.Count())

	deck = NewDeck(nil)
	for _, c := range cards {
		deck.AddCard(c)
	}

	actualCards = deck.DrawBottomN(4)
	require.Len(t, actualCards, 3)
	actualCard = actualCards[0]
	assert.Equal(t, 1, actualCard.GetValue())
	actualCard = actualCards[1]
	assert.Equal(t, 2, actualCard.GetValue())
	actualCard = actualCards[2]
	assert.Equal(t, 3, actualCard.GetValue())
	assert.Equal(t, 0, deck.Count())
}

// TODO: TestDrawCardsByIDs

// TODO: TestDrawMidN

// TODO: TestDrawRandN

func TestDrawTopN(t *testing.T) {
	cards := []Card{
		NewCard(nil, 1, ``),
		NewCard(nil, 2, ``),
		NewCard(nil, 3, ``),
	}

	deck := NewDeck(nil)
	for _, c := range cards {
		deck.AddCard(c)
	}

	actualCards := deck.DrawTopN(0)
	require.Len(t, actualCards, 0)
	assert.Equal(t, 3, deck.Count())

	deck = NewDeck(nil)
	for _, c := range cards {
		deck.AddCard(c)
	}

	actualCards = deck.DrawTopN(1)
	require.Len(t, actualCards, 1)
	actualCard := actualCards[0]
	assert.Equal(t, 1, actualCard.GetValue())
	assert.Equal(t, 2, deck.Count())

	deck = NewDeck(nil)
	for _, c := range cards {
		deck.AddCard(c)
	}

	actualCards = deck.DrawTopN(2)
	require.Len(t, actualCards, 2)
	actualCard = actualCards[0]
	assert.Equal(t, 1, actualCard.GetValue())
	actualCard = actualCards[1]
	assert.Equal(t, 2, actualCard.GetValue())
	assert.Equal(t, 1, deck.Count())

	deck = NewDeck(nil)
	for _, c := range cards {
		deck.AddCard(c)
	}

	actualCards = deck.DrawTopN(3)
	require.Len(t, actualCards, 3)
	actualCard = actualCards[0]
	assert.Equal(t, 1, actualCard.GetValue())
	actualCard = actualCards[1]
	assert.Equal(t, 2, actualCard.GetValue())
	actualCard = actualCards[2]
	assert.Equal(t, 3, actualCard.GetValue())
	assert.Equal(t, 0, deck.Count())

	deck = NewDeck(nil)
	for _, c := range cards {
		deck.AddCard(c)
	}

	actualCards = deck.DrawTopN(4)
	require.Len(t, actualCards, 3)
	actualCard = actualCards[0]
	assert.Equal(t, 1, actualCard.GetValue())
	actualCard = actualCards[1]
	assert.Equal(t, 2, actualCard.GetValue())
	actualCard = actualCards[2]
	assert.Equal(t, 3, actualCard.GetValue())
	assert.Equal(t, 0, deck.Count())
}

// TODO: TestEmpty

// TODO: TestGetBottomN

// TODO: TestGetCards

// TODO: TestGetCardsByIDs

// TODO: TestGetMidN

// TODO: TestGetRandN

// TODO: TestGetTopN

// TODO: TestHasCard

// TODO: TestIsEmpty

// TODO: TestShuffle
