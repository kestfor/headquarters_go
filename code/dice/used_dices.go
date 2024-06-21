package dice

const (
	DICE         = "🎲"
	DART         = "🎯"
	BASKETBALL   = "🏀"
	FOOTBALL     = "⚽"
	SLOT_MACHINE = "🎰"
	BOWLING      = "🎳"
)

var basketDice = Dice{wins: []int{4, 5}, Emoji: BASKETBALL}
var footballDice = Dice{wins: []int{3, 4, 5}, Emoji: FOOTBALL}
var bowlingDice = Dice{wins: []int{6}, Emoji: BOWLING}

var Collection = NewDiceCollection(bowlingDice, footballDice, basketDice)
