package tests

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"headquarters/code/dice"
	"testing"
)

func TestBasketDiceWin(t *testing.T) {
	collection := dice.Collection
	basket := tgbotapi.Dice{Emoji: dice.BASKETBALL, Value: 4}
	if !collection.Success(basket) {
		t.Error("dice should be success")
	}
}

func TestBasketDiceLose(t *testing.T) {
	collection := dice.Collection
	basket := tgbotapi.Dice{Emoji: dice.BASKETBALL, Value: 1}
	if collection.Success(basket) {
		t.Error("dice should be lose")
	}
}

func TestFootballDiceLose(t *testing.T) {
	collection := dice.Collection
	football := tgbotapi.Dice{Emoji: dice.FOOTBALL, Value: 1}
	if collection.Success(football) {
		t.Error("dice should be lose")
	}
}

func TestFootballDiceWin(t *testing.T) {
	collection := dice.Collection
	football := tgbotapi.Dice{Emoji: dice.FOOTBALL, Value: 3}
	if !collection.Success(football) {
		t.Error("dice should be win")
	}
}

func TestBowlingDiceWin(t *testing.T) {
	collection := dice.Collection
	bowling := tgbotapi.Dice{Emoji: dice.BOWLING, Value: 6}
	if !collection.Success(bowling) {
		t.Error("dice should be win")
	}
}

func TestBowlingDiceLose(t *testing.T) {
	collection := dice.Collection
	bowling := tgbotapi.Dice{Emoji: dice.BOWLING, Value: 1}
	if collection.Success(bowling) {
		t.Error("dice should be lose")
	}
}
