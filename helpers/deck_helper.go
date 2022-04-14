/*
Helper file to Create Deck and deck related methods
*/

package helpers

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pianisimo/cards-api/models"
	"math/rand"
)

var (
	ErrorEmptyDeck = errors.New("deck is empty, not possible to draw a card")
)

func NewDeck(cardCodes []string) (models.Deck, error) {
	var cards []models.Card
	var err error

	if len(cardCodes) == 0 {
		cards = make([]models.Card, 52)

		for i, suit := range Suits {
			for j, value := range Values {
				cards[j+i*13], err = CreateCardFromCode(value + suit)
			}
		}
	} else {
		cards = make([]models.Card, len(cardCodes))
		for j, code := range cardCodes {
			cards[j], err = CreateCardFromCode(code)
			if err != nil {
				return models.Deck{}, err
			}
		}
	}

	return models.Deck{
		ID:        uuid.New().String(),
		Shuffled:  false,
		Cards:     &cards,
		Remaining: uint(len(cards)),
	}, nil
}

func NewDeckFullDeck() (models.Deck, error) {
	return NewDeck([]string{})
}

func ShuffleDeck(deck *models.Deck) {
	rand.Shuffle(len(*deck.Cards), func(i, j int) {
		(*deck.Cards)[i], (*deck.Cards)[j] = (*deck.Cards)[j], (*deck.Cards)[i]
	})

	deck.Shuffled = true
}

func DrawCard(deck *models.Deck) (models.Card, error) {
	if len(*deck.Cards) < 1 {
		return models.Card{}, ErrorEmptyDeck
	}
	var popped models.Card

	popped = (*deck.Cards)[0]
	remaining := make([]models.Card, len(*deck.Cards)-1)
	copy(remaining, (*deck.Cards)[1:])
	deck.Cards = &remaining
	deck.Remaining = deck.GetRemaining()
	return popped, nil
}
