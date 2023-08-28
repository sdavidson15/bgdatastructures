package datastructures

type TurnBasedPlayer interface {
	GetName() string
	GetOrder() int32
	GetUUID() string
	IsOutOfPlay() bool
	SetOrder(int32)
	SetOutOfPlay(bool)
}

type TurnBasedPlayers []TurnBasedPlayer

type BaseTurnBasedPlayer struct {
	Structure
	Order     int32 `json:"Order"`
	OutOfPlay bool  `json:"OutOfPlay"`
}

func (b *BaseTurnBasedPlayer) GetOrder() int32 {
	return b.Order
}

func (b *BaseTurnBasedPlayer) IsOutOfPlay() bool {
	return b.OutOfPlay
}

func (b *BaseTurnBasedPlayer) SetOrder(order int32) {
	b.Order = order
}

func (b *BaseTurnBasedPlayer) SetOutOfPlay(outOfPlay bool) {
	b.OutOfPlay = outOfPlay
}

func NewTurnBasedPlayer(name *string, order *int32) TurnBasedPlayer {
	if order == nil {
		tmp := int32(-1)
		order = &tmp
	}
	return &BaseTurnBasedPlayer{
		Structure: *NewStructure(name),
		Order:     *order,
	}
}
