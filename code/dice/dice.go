package dice

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Dice struct {
	wins  []int
	Emoji string
}

type diceInterface interface {
	success(int) bool
}

func (dice *Dice) success(value int) bool {
	for _, w := range dice.wins {
		if w == value {
			return true
		}
	}
	return false
}

type DiceCollection struct {
	dices map[string]Dice
}

type DiceCollectionInterface interface {
	Success(dice tg.Dice) bool
}

func NewDiceCollection(dices ...Dice) *DiceCollection {
	collection := &DiceCollection{}
	collection.dices = make(map[string]Dice)
	for _, dice := range dices {
		collection.dices[dice.Emoji] = dice
	}
	return collection
}

func (collection *DiceCollection) Success(dice tg.Dice) bool {
	_dice, ok := collection.dices[dice.Emoji]
	if ok {
		return _dice.success(dice.Value)
	}
	return false
}
