package day7

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type CardType int

var (
	valuesPart1 = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	valuesPart2 = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
)

const (
	HighCard CardType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Bid struct {
	Hand      string
	BidAmount int
	CardType  CardType

	ModifiedHand     string
	ModifiedCardType CardType
}

func ResolveJokers(input string) (output string) {
	if !strings.Contains(input, "J") {
		// no change required
		return input
	}

	indexOf := func(v rune) int {
		for i, item := range valuesPart1 {
			if item == v {
				return i
			}
		}

		panic("not valid !!!")
	}

	values := make(map[rune]int)
	for _, r := range input {
		values[r]++
	}

	var (
		maxValue = "J"
		maxCount int
		maxIndex int
	)

	for k, v := range values {
		if k == 'J' {
			continue
		}

		ki := indexOf(k)

		if v > maxCount {
			maxCount = v
			maxValue = string(k)
			maxIndex = ki
		} else if v == maxCount {
			if ki > maxIndex {
				maxCount = v
				maxValue = string(k)
				maxIndex = ki
			}
		}
	}

	return strings.ReplaceAll(input, "J", maxValue)
}

func ResolveType(input string) CardType {
	maxValue := func(values map[rune]int) int {
		m := -1
		for _, v := range values {
			if v > m {
				m = v
			}
		}
		return m
	}

	values := make(map[rune]int)
	for _, r := range input {
		values[r]++
	}

	switch len(values) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		switch maxValue(values) {
		case 2:
			return TwoPair
		case 3:
			return ThreeOfAKind
		}
	case 2:
		switch maxValue(values) {
		case 3:
			return FullHouse
		case 4:
			return FourOfAKind
		}
	case 1:
		return FiveOfAKind
	}

	log.Println(input, values)
	panic("unreachable !!!")
}

func ComparePart1(a, b Bid) int {
	indexOf := func(v rune) int {
		for i, item := range valuesPart1 {
			if item == v {
				return i
			}
		}

		panic("not valid !!!")
	}

	if a.CardType != b.CardType {
		if a.CardType > b.CardType {
			return 1
		}

		return -1
	}

	for i := 0; i < 5; i++ {
		av := rune(a.Hand[i])
		bv := rune(b.Hand[i])

		if av == bv {
			continue
		}

		ai := indexOf(av)
		bi := indexOf(bv)
		if ai > bi {
			return 1
		}

		return -1
	}

	return 0
}

func ComparePart2(a, b Bid) int {
	indexOf := func(v rune) int {
		for i, item := range valuesPart2 {
			if item == v {
				return i
			}
		}

		panic("not valid !!!")
	}

	if a.ModifiedCardType != b.ModifiedCardType {
		if a.ModifiedCardType > b.ModifiedCardType {
			return 1
		}

		return -1
	}

	for i := 0; i < 5; i++ {
		av := rune(a.Hand[i])
		bv := rune(b.Hand[i])

		if av == bv {
			continue
		}

		ai := indexOf(av)
		bi := indexOf(bv)
		if ai > bi {
			return 1
		}

		return -1
	}

	return 0
}

func ReadBids(input string) []Bid {
	var (
		bids = make([]Bid, 0)
	)

	for i, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bidAmount, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(i, line, err)
		}

		b := Bid{
			Hand:      hand,
			BidAmount: bidAmount,
			CardType:  ResolveType(hand),
		}

		bids = append(bids, b)
	}

	return bids
}

func Part1(bids []Bid) int {
	slices.SortFunc(bids, ComparePart1)
	amount := 0

	log.Println(bids)

	for i, b := range bids {
		amount += (i + 1) * b.BidAmount
	}

	return amount
}

func Part2(bids []Bid) int {
	for i, b := range bids {
		bids[i].ModifiedHand = ResolveJokers(b.Hand)
		bids[i].ModifiedCardType = ResolveType(bids[i].ModifiedHand)
	}

	slices.SortFunc(bids, ComparePart2)
	amount := 0

	log.Println(bids)

	for i, b := range bids {
		amount += (i + 1) * b.BidAmount
	}

	return amount
}
