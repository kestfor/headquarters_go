package utils

import "math/rand"

func GetRandomChallengeEmoji() string {
	emojis := []rune(`😐🤭🙄😅🥲😙🤨🤪🥸`)
	maxN := len(emojis)
	return string(emojis[rand.Intn(maxN)])
}

func GetRandomHappyEmoji() string {
	emojis := []rune(`🥹😌😉😜😎🤩🥳😏🤗🫡`)
	maxN := len(emojis)
	return string(emojis[rand.Intn(maxN)])
}
