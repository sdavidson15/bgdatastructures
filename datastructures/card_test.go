package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	expectedValue := 10
	card := NewCard(nil, expectedValue, `testSuit`)
	assert.Equal(t, expectedValue, card.GetValue())
}

func TestGetSuit(t *testing.T) {
	expectedSuit := `testSuit`
	card := NewCard(nil, -1, expectedSuit)
	assert.Equal(t, expectedSuit, card.GetSuit())
}

func TestNewCard(t *testing.T) {
	expectedName := `testName`
	expectedValue := 10
	expectedSuit := `testSuit`

	card := NewCard(nil, expectedValue, expectedSuit)
	assert.Greater(t, len(card.GetUUID()), 0)
	assert.Equal(t, `<nil>`, card.GetName())
	assert.Equal(t, expectedValue, card.GetValue())
	assert.Equal(t, expectedSuit, card.GetSuit())

	card = NewCard(&expectedName, expectedValue, expectedSuit)
	assert.Greater(t, len(card.GetUUID()), 0)
	assert.Equal(t, expectedName, card.GetName())
	assert.Equal(t, expectedValue, card.GetValue())
	assert.Equal(t, expectedSuit, card.GetSuit())
}
