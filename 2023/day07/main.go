package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"adtennant.dev/aoc/util"
)

type hand struct {
	cards []byte
	bid   int
}

type HandKind int

const (
	HAND_KIND_FIVE_OF_A_KIND HandKind = iota
	HAND_KIND_FOUR_OF_A_KIND
	HAND_KIND_FULL_HOUSE
	HAND_KIND_THREE_OF_A_KIND
	HAND_KIND_TWO_PAIR
	HAND_KIND_PAIR
	HAND_KIND_HIGH_CARD
)

func (h hand) kind(useJokers bool) HandKind {
	cards := make(map[byte]int)

	for _, c := range h.cards {
		cards[c] += 1
	}

	if useJokers {
		jokers := cards['J']
		delete(cards, 'J')

		maxCard := byte('X')
		maxCount := -1

		for card, count := range cards {
			if count > maxCount {
				maxCard = card
				maxCount = count
			}
		}

		cards[maxCard] += jokers
	}

	groups := make(map[int]int)

	for _, s := range cards {
		groups[s] += 1
	}

	switch {
	case groups[5] == 1:
		return HAND_KIND_FIVE_OF_A_KIND
	case groups[4] == 1:
		return HAND_KIND_FOUR_OF_A_KIND
	case groups[3] == 1 && groups[2] == 1:
		return HAND_KIND_FULL_HOUSE
	case groups[3] == 1:
		return HAND_KIND_THREE_OF_A_KIND
	case groups[2] == 2:
		return HAND_KIND_TWO_PAIR
	case groups[2] == 1:
		return HAND_KIND_PAIR
	default:
		return HAND_KIND_HIGH_CARD
	}
}

func parseHand(str string) (hand, error) {
	parts := strings.Split(str, " ")
	if len(parts) != 2 {
		return hand{}, fmt.Errorf("invalid hand format")
	}

	cards := []byte(parts[0])
	if len(cards) != 5 {
		return hand{}, fmt.Errorf("invalid number of cards")
	}

	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		return hand{}, err
	}

	return hand{
		cards,
		bid,
	}, nil
}

func parseHands(input string) (hands []hand, err error) {
	for _, line := range util.Lines(input) {
		hand, err := parseHand(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse hand: %w", err)
		}

		hands = append(hands, hand)
	}

	return hands, nil
}

func calculateWinnings(hands []hand, cardRanks []byte, useJokers bool) int {
	slices.SortFunc(hands, func(a, b hand) int {
		kind := b.kind(useJokers) - a.kind(useJokers)

		if kind != 0 {
			return int(kind)
		}

		for i := 0; i < 5; i++ {
			cardA := a.cards[i]
			cardB := b.cards[i]

			if cardA != cardB {
				return slices.Index(cardRanks, cardB) - slices.Index(cardRanks, cardA)
			}
		}

		return 0
	})

	winnings := 0

	for i, h := range hands {
		winnings += (i + 1) * h.bid
	}

	return winnings
}

func Part1(input string) (int, error) {
	hands, err := parseHands(input)
	if err != nil {
		return -1, err
	}

	return calculateWinnings(hands, []byte{
		'A',
		'K',
		'Q',
		'J',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
	}, false), nil
}

func Part2(input string) (int, error) {
	hands, err := parseHands(input)
	if err != nil {
		return -1, err
	}

	return calculateWinnings(hands, []byte{
		'A',
		'K',
		'Q',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
		'J',
	}, true), nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
