package parser_test

import (
	"fmt"
	"testing"

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

	tags := []string{"y", "z", "x", "0", "1", "2"}
	res := parser.MainTrd(trades, tags)

	// iterate over every key to print the result
	for k, v := range res {
		fmt.Printf("Key: %v; Trades: %v", k, v)
		fmt.Println()
	}
}
