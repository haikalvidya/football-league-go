package usecase

import (
	"strings"

	"github.com/haikalvidya/football-league/internal/delivery/payload"
)

type IContainLetterUsecase interface {
	ContainLetter(*payload.ContainLetterRequest) bool
}

type containLetterUsecase usecaseType

func (u *containLetterUsecase) ContainLetter(req *payload.ContainLetterRequest) bool {
	// to lower
	req.FirstWord = strings.ToLower(req.FirstWord)
	req.SecondWord = strings.ToLower(req.SecondWord)
	charCounts := make(map[rune]int)
	for _, char := range req.SecondWord {
		charCounts[char]++
	}
	for _, char := range req.FirstWord {
		if charCounts[char] == 0 {
			return false
		}
		charCounts[char]--
	}
	return true

}
