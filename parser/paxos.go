package parser

// Trade:
// id, type (BUY/SELL), totalGold, totalUSD
// where:
// totalGold, totalUSD > 0

// Net Position:
// netGold, netUSD
// where:
// -Infinity < netGold, netUSD < Infinity

// Problem 1:
// Input: List of `Trade`s
// Output: Net Position for those trades
// Example:
// For the two trades with ids “A” and “B” here

//netP = [ 0, 0 ]
//  * “A”, BUY, totalGold = 1, totalUSD = 1750
//netP = [ 1, -1750 ]
//  * “B”, SELL, totalGold = 0.25, totalUSD = 750
//netP = [ 0.75, -1000 ]
// The Net Position is:
//     netGold = 0.75, netUSD = -1000

// Trade:
// id, type (BUY/SELL), totalGold, totalUSD, ourTag, theirTag
// where:
// totalGold, totalUSD > 0
// for all ourTag theirTag, ourTag != theirTag (these are non-intersecting sets)
// Trades:

// * A: ourTag = "x", theirTag = "0"
// * B: ourTag = "x", theirTag = "1"
// * C: ourTag = "y", theirTag = "1"
// * D: ourTag = "z", theirTag = "2"

/*x -> [ A, B ]
y -> [ C ]
z -> [ D ]
0 -> [ A ]
1 -> [ B, C ]
2 -> [ D ]

O(NLogN)*/

// A,B,C are in the same group.
// D is in a group by itself.
// Problem 2:
// Input: List of trades, a tag
// Output: The net position for the group of trades containing that tag.
//tag: 1

//1. tags = ["1"]
//2. tags = ["1", "x", "y"]
//3. tags = ["1", "x", "y", "0"]

type NetPosition struct {
	Gold float64
	USD  float64
}

type Tag string

type Trade struct {
	ID        int
	Type      string
	TotalUSD  float64
	TotalGold float64
	OutTag    Tag
	TheirTag  Tag
}

func getNetPosition(trades []Trade) NetPosition {
	var netPosition NetPosition
	for _, trade := range trades {
		if trade.Type == "BUY" {
			netPosition.Gold = netPosition.Gold + trade.TotalGold
			netPosition.USD = netPosition.USD - trade.TotalUSD
		}
		if trade.Type == "SELL" {
			netPosition.Gold = netPosition.Gold - trade.TotalGold
			netPosition.USD = netPosition.USD + trade.TotalUSD
		}
	}

	return netPosition
}

func NetPositionByTag(trades []Trade, tag Tag) NetPosition {
	var tagTrades []Trade
	for _, t := range trades {
		if t.OutTag == tag {
			tagTrades = append(tagTrades, t)
		}
		if t.TheirTag == tag {
			tagTrades = append(tagTrades, t)
		}
	}

	return getNetPosition(tagTrades)
}

func MainNetPosition(trades []Trade, tags []Tag) map[Tag]NetPosition {
	tt := map[Tag]NetPosition{}
	for _, t := range tags {
		nPos := NetPositionByTag(trades, t)
		tt[t] = nPos
	}
	return tt
}
