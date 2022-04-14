package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pianisimo/cards-api/database"
	"github.com/pianisimo/cards-api/models"
)

var (
	validate                       = validator.New()
	repository database.Repository = &database.MyMemoryRepository
	// TODO
	/*
		implement more repositories, like mongoDB, firestore, etc.
		make sure all the repositories implements the database.Repository Interface
	*/
)

func CreateNewDeck(ctx *fiber.Ctx) error {
	deckRequest := new(models.DeckRequest)

	err := ctx.BodyParser(deckRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	newDeck, err := database.MyMemoryRepository.CreateNewDeck(deckRequest.GetCardCodesFromString(deckRequest.Cards), deckRequest.Shuffled)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(newDeck)
}

func DrawCards(ctx *fiber.Ctx) error {
	drawCardRequest := new(models.DrawCardRequest)
	err := ctx.BodyParser(drawCardRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err = validate.Struct(drawCardRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	cards, err := database.MyMemoryRepository.DrawCards(drawCardRequest.Id, drawCardRequest.Number)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(cards)
}

func OpenDeck(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	newDeck, err := database.MyMemoryRepository.OpenDeck(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(200).JSON(newDeck)
}

func GetAllDecks(ctx *fiber.Ctx) error {
	decks := database.MyMemoryRepository.GetAllDecks()
	return ctx.Status(200).JSON(decks)
}
