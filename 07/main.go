package main

import (
	"aoc2023/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand string
	bid  int
}

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(input []byte) int {
	return getTotalWinnings(input, false)
}

func partTwo(input []byte) int {
	return getTotalWinnings(input, true)
}

func getTotalWinnings(input []byte, withJoker bool) int {
	hands := transformInput(input)

	sortHands(hands, withJoker)

	totalWinnings := 0

	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}

	return totalWinnings
}

func transformInput(input []byte) []Hand {
	lines := strings.Split(string(input), "\n")

	hands := make([]Hand, 0, 100)

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		bid, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic(fmt.Sprintf("Bid needs to be a valid integer: %s", err))
		}

		hands = append(hands, Hand{hand: splitLine[0], bid: bid})
	}

	return hands
}

func sortHands(hands []Hand, withJoker bool) {
	sort.Slice(hands, func(a, b int) bool {
		handA := hands[a].hand
		handB := hands[b].hand

		var valueA, valueB int
		if withJoker {
			valueA = determineValue(replaceJoker(handA))
			valueB = determineValue(replaceJoker(handB))
		} else {
			valueA = determineValue(handA)
			valueB = determineValue(handB)
		}

		if valueA == valueB {
			for i := 0; i < len(handA); i++ {
				if handA[i] == handB[i] {
					continue
				}

				valueA := convertCardToNumber(handA[i], withJoker)
				valueB := convertCardToNumber(handB[i], withJoker)

				if valueA == valueB {
					continue
				}

				return valueA < valueB

			}
		}

		return valueA < valueB
	})
}

func convertCardToNumber(card byte, weakJoker bool) int {
	if card >= '2' && card <= '9' {
		return int(card - '0')
	}

	switch card {
	case 'T':
		return 10
	case 'J':
		if weakJoker {
			return -1
		}
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		panic(fmt.Sprintf("Unexpected card: %s", string(card)))
	}
}

func determineValue(hand string) int {
	cardCount := make(map[rune]int, 5)

	for _, card := range hand {
		cardCount[card] = cardCount[card] + 1
	}

	if len(cardCount) == 5 {
		return 0
	}

	if len(cardCount) == 4 {
		return 1
	}

	if len(cardCount) == 3 {
		for _, count := range cardCount {
			if count == 3 {
				return 3
			}
		}
		return 2
	}

	if len(cardCount) == 2 {
		for _, count := range cardCount {
			if count == 3 {
				return 4
			}

			if count == 4 {
				return 5
			}
		}
		return 3
	}

	return 6
}

func replaceJoker(hand string) string {
	cardCount := make(map[rune]int, 5)

	for _, card := range hand {
		cardCount[card] = cardCount[card] + 1
	}

	higestCardCount := -1
	highestCard := '2'
	numberOfJokers := 0

	for card, count := range cardCount {
		if card == 'J' {
			numberOfJokers++
		} else if count > higestCardCount {
			higestCardCount = count
			highestCard = card
		} else if count == higestCardCount {
			if convertCardToNumber(byte(card), true) > convertCardToNumber(byte(highestCard), true) {
				higestCardCount = count
				highestCard = card
			}
		}
	}

	if numberOfJokers != 0 {
		cardCount[highestCard] = cardCount[highestCard] + numberOfJokers
		hand = strings.ReplaceAll(hand, "J", string(highestCard))
		delete(cardCount, 'J')
	}

	return hand
}
