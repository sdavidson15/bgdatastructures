package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// FIXME: use a mock player?

func TestNextTurn(t *testing.T) {
	g := &BaseTurnBasedGame{
		Structure: *NewStructure(nil),
		Order:     -1,
		Players:   TurnBasedPlayers{},
	}

	g.NextTurn()
	assert.Equal(t, int32(-1), g.GetCurrentOrder())

	fstPlayerName := `edith`
	fstPlayer := NewTurnBasedPlayer(&fstPlayerName, nil)
	_, err := g.AddPlayer(fstPlayer)
	require.NoError(t, err)
	require.Equal(t, int32(0), fstPlayer.GetOrder())
	require.Equal(t, int32(0), g.GetCurrentOrder())

	g.NextTurn()
	require.Equal(t, int32(0), g.GetCurrentOrder())

	sndPlayerName := `naveen`
	sndPlayer := NewTurnBasedPlayer(&sndPlayerName, nil)
	_, err = g.AddPlayer(sndPlayer)
	require.NoError(t, err)
	require.Equal(t, int32(1), sndPlayer.GetOrder())
	require.Equal(t, int32(0), g.GetCurrentOrder())

	g.NextTurn()
	require.Equal(t, int32(1), g.GetCurrentOrder())

	thdPlayerName := `naveen`
	thdPlayerOrder := int32(5)
	thdPlayer := NewTurnBasedPlayer(&thdPlayerName, &thdPlayerOrder)
	_, err = g.AddPlayer(thdPlayer)
	require.NoError(t, err)
	require.Equal(t, int32(5), thdPlayer.GetOrder())
	require.Equal(t, int32(1), g.GetCurrentOrder())

	g.NextTurn()
	require.Equal(t, int32(5), g.GetCurrentOrder())

	g.NextTurn()
	require.Equal(t, int32(0), g.GetCurrentOrder())

	g.NextTurn()
	require.Equal(t, int32(1), g.GetCurrentOrder())

	g.NextTurn()
	require.Equal(t, int32(5), g.GetCurrentOrder())

	g.NextTurn()
	require.Equal(t, int32(0), g.GetCurrentOrder())
}
