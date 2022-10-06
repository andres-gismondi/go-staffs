package parser_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-staffs/parser"
)

func TestMindera_IsPaxos_NotMindera(t *testing.T) {
	trades := []parser.Trd{
		{
			ID:     "A",
			Type:   "BUY",
			Gold:   1,
			USD:    1000,
			OurT:   "x",
			TheirT: "0",
		},
		{
			ID:     "B",
			Type:   "SELL",
			Gold:   0.25,
			USD:    750,
			OurT:   "x",
			TheirT: "1",
		},
		{
			ID:     "C",
			Type:   "BUY",
			Gold:   2,
			USD:    2000,
			OurT:   "y",
			TheirT: "1",
		},
		{
			ID:     "D",
			Type:   "SELL",
			Gold:   1,
			USD:    1500,
			OurT:   "z",
			TheirT: "2",
		},
	}

	got := parser.MainTrd(trades, "x")
	fmt.Printf("Trades: %v", got)
	assert.Equal(t, 3, len(got))

	fmt.Println()
	got = parser.MainTrd(trades, "1")
	fmt.Printf("Trades: %v", got)
	assert.Equal(t, 3, len(got))

	fmt.Println()
	got = parser.MainTrd(trades, "y")
	fmt.Printf("Trades: %v", got)
	assert.Equal(t, 3, len(got))

	fmt.Println()
	got = parser.MainTrd(trades, "0")
	fmt.Printf("Trades: %v", got)
	assert.Equal(t, 3, len(got))

	fmt.Println()
	got = parser.MainTrd(trades, "z")
	fmt.Printf("Trades: %v", got)
	assert.Equal(t, 1, len(got))
}
