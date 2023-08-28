package datastructures

import "math/rand"

type Deck interface {
	AddCard(Card)
	Count() int
	DrawBottomN(int) Cards
	DrawCardsByIDs(...string) Cards
	DrawMidN(int, int) Cards
	DrawRandN(int) Cards
	DrawTopN(int) Cards
	Empty() Cards
	GetBottomN(int) Cards
	GetCards() Cards
	GetCardsByIDs(...string) Cards
	GetMidN(int, int) Cards
	GetRandN(int) Cards
	GetTopN(int) Cards
	GetUUID() string
	GetName() string
	HasCard(Card, func(Card) bool) bool
	IsEmpty() bool
	Shuffle()
}

type Decks []Deck

type BaseDeck struct {
	Structure
	Cards Cards
}

func (b *BaseDeck) AddCard(cards Card) {
	b.Cards = append(b.Cards, cards)
}

func (b *BaseDeck) Count() int {
	return len(b.Cards)
}

func (b *BaseDeck) DrawBottomN(n int) Cards {
	n = min(len(b.Cards), n)
	drawn := b.Cards[len(b.Cards)-n:]
	b.Cards = b.Cards[:len(b.Cards)-n]
	return drawn
}

func (b *BaseDeck) DrawCardsByIDs(ids ...string) Cards {
	drawn := make(Cards, 0, min(len(ids), len(b.Cards)))
	notDrawn := make(Cards, 0, len(b.Cards))
	idsMap := map[string]struct{}{}
	for _, id := range ids {
		idsMap[id] = struct{}{}
	}
	for _, c := range b.Cards {
		if _, ok := idsMap[c.GetUUID()]; ok {
			drawn = append(drawn, c)
		} else {
			notDrawn = append(notDrawn, c)
		}
	}
	b.Cards = notDrawn
	return drawn
}

func (b *BaseDeck) DrawMidN(offset int, n int) Cards {
	// TODO: stub
	return nil
}

func (b *BaseDeck) DrawRandN(n int) Cards {
	// TODO: stub
	return nil
}

func (b *BaseDeck) DrawTopN(n int) Cards {
	n = min(len(b.Cards), n)
	drawn := b.Cards[:n]
	b.Cards = b.Cards[n:]
	return drawn
}

func (b *BaseDeck) Empty() Cards {
	return b.DrawTopN(b.Count())
}

func (b *BaseDeck) GetBottomN(n int) Cards {
	n = min(len(b.Cards), n)
	return b.Cards[len(b.Cards)-n:]
}

func (b *BaseDeck) GetCards() Cards {
	return b.Cards
}

func (b *BaseDeck) GetCardsByIDs(ids ...string) Cards {
	drawn := make(Cards, 0, min(len(ids), len(b.Cards)))
	idsMap := map[string]struct{}{}
	for _, id := range ids {
		idsMap[id] = struct{}{}
	}
	for _, c := range b.Cards {
		if _, ok := idsMap[c.GetUUID()]; ok {
			drawn = append(drawn, c)
		}
	}
	return drawn
}

func (b *BaseDeck) GetMidN(offset int, n int) Cards {
	// TODO: stub
	return nil
}

func (b *BaseDeck) GetRandN(n int) Cards {
	if len(b.Cards) == 0 {
		return b.Cards
	}

	n = min(len(b.Cards), n)
	indices := rand.Perm(len(b.Cards))
	selected := make([]Card, n)
	for i := 0; i < n; i++ {
		selected[i] = b.Cards[indices[i]]
	}
	return selected
}

func (b *BaseDeck) GetTopN(n int) Cards {
	n = min(len(b.Cards), n)
	return b.Cards[:n]
}

func (b *BaseDeck) HasCard(card Card, matcher func(Card) bool) bool {
	for _, c := range b.Cards {
		if matcher(c) {
			return true
		}
	}
	return false
}

func (b *BaseDeck) IsEmpty() bool {
	return len(b.Cards) == 0
}

func (b *BaseDeck) Shuffle() {
	rand.Shuffle(len(b.Cards), func(i int, j int) {
		b.Cards[i], b.Cards[j] = b.Cards[j], b.Cards[i]
	})
}

func NewDeck(name *string) Deck {
	return &BaseDeck{
		Structure: *NewStructure(name),
		Cards:     Cards{},
	}
}

func CardsToDeck(cards Cards) Deck {
	deck := NewDeck(nil)
	for _, c := range cards {
		deck.AddCard(c)
	}
	return deck
}
