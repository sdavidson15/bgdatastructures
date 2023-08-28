package datastructures

import (
	"fmt"
)

type TurnBasedGame interface {
	AddPlayer(TurnBasedPlayer) (TurnBasedPlayer, error)
	GetCurrentOrder() int32
	GetCurrentRound() int32
	GetCurrentPlayer() TurnBasedPlayer
	GetCurrentPlayerID() string
	GetPlayerByID(string) TurnBasedPlayer
	GetPlayers() TurnBasedPlayers
	GetName() string
	GetUUID() string
	MaxOrder() int32
	NextTurn() int32
	RemovePlayer(string) TurnBasedPlayer
	StartGame() error
}

type BaseTurnBasedGame struct {
	Structure
	Order   int32            `json:"Order"`
	Round   int32            `json:"Round"`
	Players TurnBasedPlayers `json:"Players,omitempty"`
}

func (b *BaseTurnBasedGame) AddPlayer(player TurnBasedPlayer) (TurnBasedPlayer, error) {
	for _, p := range b.Players {
		if p.GetOrder() == player.GetOrder() {
			return nil, fmt.Errorf(`cannot add player to game, order must be unique`)
		}
	}

	if player.GetOrder() == -1 {
		player.SetOrder(b.MaxOrder() + 1)
	}

	if b.GetCurrentOrder() == -1 {
		b.Order = player.GetOrder()
	}

	b.Players = append(b.Players, player)
	return player, nil
}

func (b *BaseTurnBasedGame) GetCurrentOrder() int32 {
	return b.Order
}

func (b *BaseTurnBasedGame) GetCurrentRound() int32 {
	return b.Round
}

func (b *BaseTurnBasedGame) GetCurrentPlayer() TurnBasedPlayer {
	for _, p := range b.Players {
		if p.GetOrder() == b.Order {
			return p
		}
	}
	return nil
}

func (b *BaseTurnBasedGame) GetCurrentPlayerID() string {
	return b.GetCurrentPlayer().GetUUID()
}

func (b *BaseTurnBasedGame) GetPlayerByID(id string) TurnBasedPlayer {
	for _, p := range b.Players {
		if p.GetUUID() == id {
			return p
		}
	}
	return nil
}

func (b *BaseTurnBasedGame) GetPlayers() TurnBasedPlayers {
	return b.Players
}

func (b *BaseTurnBasedGame) MaxOrder() int32 {
	max := b.Order
	for _, p := range b.Players {
		if !p.IsOutOfPlay() && p.GetOrder() > max {
			max = p.GetOrder()
		}
	}
	return max
}

func (b *BaseTurnBasedGame) NextTurn() int32 {
	var min, next int32 = b.MaxOrder(), b.Order
	for _, p := range b.Players {
		if !p.IsOutOfPlay() {
			pOrder := p.GetOrder()
			if pOrder > b.Order {
				if next == b.Order {
					next = pOrder
				} else if pOrder < next {
					next = pOrder
				}
			}
			if pOrder < min {
				min = pOrder
			}
		}
	}

	if next == b.Order {
		b.Round += 1
		next = min
	}
	b.Order = next
	return b.Order
}

func (b *BaseTurnBasedGame) RemovePlayer(id string) TurnBasedPlayer {
	var removed TurnBasedPlayer
	filtered := make([]TurnBasedPlayer, 0, len(b.Players))
	for _, p := range b.Players {
		if p.GetUUID() != id {
			filtered = append(filtered, p)
		} else {
			removed = p
		}
	}
	b.Players = filtered
	if removed.GetOrder() == b.Order {
		b.NextTurn()
	}
	return removed
}

func (b *BaseTurnBasedGame) StartGame() error {
	if b.GetCurrentOrder() < 0 || len(b.Players) == 0 {
		return fmt.Errorf(`cannot start a game with no players`)
	}
	if b.GetCurrentRound() > 0 {
		return fmt.Errorf(`cannot start a game that has already started`)
	}
	b.Round = 0
	return nil
}

func NewTurnBasedGame(name *string) TurnBasedGame {
	return &BaseTurnBasedGame{
		Structure: *NewStructure(name),
		Order:     -1,
		Round:     0,
		Players:   TurnBasedPlayers{},
	}
}
