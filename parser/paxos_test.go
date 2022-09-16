package parser_test

import (
	"fmt"
	"testing"

	"go-staffs/parser"
)

func TestPaxos_Main(t *testing.T) {
	trades := []parser.Trade{
		{
			ID:        1,
			Type:      "BUY",
			TotalGold: 1,
			TotalUSD:  1000,
			OutTag:    "a",
			TheirTag:  "1",
		},
		{
			ID:        2,
			Type:      "SELL",
			TotalGold: 0.25,
			TotalUSD:  750,
			OutTag:    "a",
			TheirTag:  "2",
		},
		{
			ID:        3,
			Type:      "BUY",
			TotalGold: 2,
			TotalUSD:  2000,
			OutTag:    "b",
			TheirTag:  "2",
		},
		{
			ID:        4,
			Type:      "SELL",
			TotalGold: 1,
			TotalUSD:  1500,
			OutTag:    "c",
			TheirTag:  "1",
		},
	}

	tags := []parser.Tag{"a", "b", "c", "1", "2"}
	m := parser.MainNetPosition(trades, tags)

	fmt.Println(m)
}
