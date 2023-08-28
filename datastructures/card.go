package datastructures

type Card interface {
	GetValue() int
	GetSuit() string
	GetUUID() string
	GetName() string
}

type Cards []Card

type BaseCard struct {
	Structure
	Value int    `json:"Value,omitempty"`
	Suit  string `json:"Suit,omitempty"`
}

func (bc *BaseCard) GetValue() int {
	return bc.Value
}

func (bc *BaseCard) GetSuit() string {
	return bc.Suit
}

func (cards Cards) GetUUIDs() []string {
	uuids := make([]string, len(cards))
	for i, c := range cards {
		uuids[i] = c.GetUUID()
	}
	return uuids
}

func NewCard(name *string, value int, suit string) Card {
	return &BaseCard{
		Structure: *NewStructure(name),
		Value:     value,
		Suit:      suit,
	}
}
