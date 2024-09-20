package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDetermineValue_HighCard(t *testing.T) {
	hand := "23456"

	value := determineValue(hand)

	assert.Equal(t, value, 0)
}

func TestDetermineValue_OnePair(t *testing.T) {
	hand := "23446"

	value := determineValue(hand)

	assert.Equal(t, value, 1)
}

func TestDetermineValue_TwoPair(t *testing.T) {
	hand := "22446"

	value := determineValue(hand)

	assert.Equal(t, value, 2)
}

func TestDetermineValue_ThreeOfAKind(t *testing.T) {
	hand := "23444"

	value := determineValue(hand)

	assert.Equal(t, value, 3)
}

func TestDetermineValue_FullHouse(t *testing.T) {
	hand := "22333"

	value := determineValue(hand)

	assert.Equal(t, value, 4)
}

func TestDetermineValue_FourOfAKind(t *testing.T) {
	hand := "2AAAA"

	value := determineValue(hand)

	assert.Equal(t, value, 5)
}

func TestDetermineValue_FiveOfAKind(t *testing.T) {
	hand := "AAAAA"

	value := determineValue(hand)

	assert.Equal(t, value, 6)
}
