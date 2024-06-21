package utils

import "math/rand"

func GetRandomChallengePhrase(phrases []string) string {
	maxN := len(phrases)
	return phrases[rand.Intn(maxN)] + GetRandomChallengeEmoji()
}
