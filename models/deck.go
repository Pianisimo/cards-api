package models

import (
	"strings"
)

type Deck struct {
	ID        string  `json:"deck_id"`
	Shuffled  bool    `json:"shuffled"`
	Cards     *[]Card `json:"cards"`
	Remaining uint    `json:"remaining"`
}

func (d Deck) GetRemaining() uint {
	return uint(len(*d.Cards))
}

type DeckRequest struct {
	Shuffled bool   `json:"shuffled"`
	Cards    string `json:"cards"`
}

type DrawCardRequest struct {
	Id     string `json:"id" validate:"required"`
	Number uint   `json:"number" validate:"required"`
}

func (d DeckRequest) GetCardCodesFromString(cardsRequest string) []string {
	split := strings.Split(cardsRequest, ",")
	if len(split) == 1 && split[0] == "" {
		split = []string{}
	}
	return split
}
