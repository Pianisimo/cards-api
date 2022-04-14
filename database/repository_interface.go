/*
All Databases clases should implement this interface.
Later you can decide in routes.public which database to use
*/

package database

import "github.com/pianisimo/cards-api/models"

type Repository interface {
	CreateNewDeck(cards []string, shuffled bool) (models.Deck, error)
	DrawCards(deckId string, cardsToDraw uint) ([]models.Card, error)
	OpenDeck(deckId string) (*models.Deck, error)
}
