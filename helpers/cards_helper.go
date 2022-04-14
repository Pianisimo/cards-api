/*
Helper file to Create models.Card
*/

package helpers

import (
	"errors"
	"github.com/pianisimo/cards-api/models"
)

var (
	Values = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	Suits  = []string{"C", "D", "H", "S"}
)

var (
	ErrorInvalidCodeFormat = errors.New("invalid card code format: expected card value as first or second character and card suit for the next character")
	ErrorInvalidCardValue  = errors.New("invalid card value: values are restricted to A,2,3,4,5,6,7,8,9,10,J,Q or K")
	ErrorInvalidCardSuit   = errors.New("invalid card suit: suits are restricted to C,D,H or S")
)

func CreateCardFromCode(code string) (models.Card, error) {
	if len(code) != 2 && len(code) != 3 {
		return models.Card{}, ErrorInvalidCodeFormat
	}

	v := code[0:1] // holding card value
	s := code[1:2] // holding card suit

	// If value is 10 the suit comes from the next position
	if v == "1" && s == "0" {
		v = "10"
		s = code[2:3]
	}

	var value string
	var suit string

	switch v {
	case "A":
		value = "ACE"
	case "2":
		value = "2"
	case "3":
		value = "3"
	case "4":
		value = "4"
	case "5":
		value = "5"
	case "6":
		value = "6"
	case "7":
		value = "7"
	case "8":
		value = "8"
	case "9":
		value = "9"
	case "10":
		value = "10"
	case "J":
		value = "JOKER"
	case "Q":
		value = "QUEEN"
	case "K":
		value = "KING"
	default:
		return models.Card{}, ErrorInvalidCardValue
	}

	switch s {
	case "C":
		suit = "CLUBS"
	case "D":
		suit = "DIAMOND"
	case "H":
		suit = "HEARTS"
	case "S":
		suit = "SPADES"
	default:
		return models.Card{}, ErrorInvalidCardSuit
	}

	return models.Card{
		Value: value,
		Suit:  suit,
		Code:  code,
	}, nil
}
