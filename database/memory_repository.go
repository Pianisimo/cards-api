/*
Using Memory as a database, not recommended.
This file is just a demo to illustrate how the database.Repository interface
should be implemented
*/

package database

import (
	"errors"
	"github.com/pianisimo/cards-api/helpers"
	"github.com/pianisimo/cards-api/models"
	"sync"
)

var (
	ErrorDeckIdNotFound = errors.New("deck with requested id not found")
	ErrorNotEnoughCards = errors.New("trying to draw more cards than deck currently has")
	mutex               = sync.RWMutex{}
)

type MemoryRepository struct {
	decks *[]models.Deck
}

var MyMemoryRepository = NewMemoryRepository()

func (m *MemoryRepository) CreateNewDeck(cards []string, shuffled bool) (models.Deck, error) {
	var deck models.Deck
	var err error

	if len(cards) == 0 {
		deck, err = helpers.NewDeckFullDeck()
	} else {
		deck, err = helpers.NewDeck(cards)
	}

	if err != nil {
		return models.Deck{}, err
	}

	if shuffled {
		helpers.ShuffleDeck(&deck)
	}

	mutex.Lock()
	*m.decks = append(*m.decks, deck)
	mutex.Unlock()
	return deck, nil
}

func (m *MemoryRepository) DrawCards(deckId string, cardsToDraw uint) ([]models.Card, error) {
	deck, err := m.OpenDeck(deckId)
	if err != nil {
		return []models.Card{}, err
	}

	if deck.GetRemaining() < cardsToDraw {
		return []models.Card{}, ErrorNotEnoughCards
	}

	var cards = make([]models.Card, cardsToDraw)
	mutex.Lock()
	for i := 0; i < int(cardsToDraw); i++ {
		cards[i], err = helpers.DrawCard(deck)
		if err != nil {
			mutex.Unlock()
			return []models.Card{}, err
		}
	}

	mutex.Unlock()
	return cards, nil
}

func (m *MemoryRepository) OpenDeck(deckId string) (*models.Deck, error) {
	mutex.RLock()
	deck := &models.Deck{}
	err := ErrorDeckIdNotFound

	for i := 0; i < len(*m.decks); i++ {
		if (*m.decks)[i].ID == deckId {
			deck = &(*m.decks)[i]
			err = nil
			mutex.RUnlock()
			return deck, err
		}
	}
	mutex.RUnlock()
	return deck, err
}

func (m *MemoryRepository) GetAllDecks() *[]models.Deck {
	return m.decks
}

func NewMemoryRepository() MemoryRepository {
	return MemoryRepository{decks: &[]models.Deck{}}
}
