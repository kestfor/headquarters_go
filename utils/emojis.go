package utils

import (
	"math/rand/v2"
	"strconv"
)

func GetRandomChallengeEmoji() string {
	emojis := "😐🤭🙄😅🥲😙🤨🤪🥸"
	maxN := len(emojis)
	return strconv.Itoa(int(emojis[rand.IntN(maxN)]))
}

func GetRandomHappyEmoji() string {
	emojis := "🥹😌😉😜😎🤩🥳😏🤗🫡"
	maxN := len(emojis)
	return strconv.Itoa(int(emojis[rand.IntN(maxN)]))
}
