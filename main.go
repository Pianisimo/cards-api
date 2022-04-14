package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/pianisimo/cards-api/routes"
	"log"
	"os"
)

func main() {
	initApp()
}

func initApp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error while loading .env file: %v", err)
	}

	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("ADDRESS") + os.Getenv("PORT")))
}

func setupRoutes(app *fiber.App) {
	app.Post("/create-new-deck", routes.CreateNewDeck)
	app.Put("/draw-cards", routes.DrawCards)
	app.Get("/open-deck/:id", routes.OpenDeck)
	app.Get("/get-all-decks", routes.GetAllDecks)
	// TODO
	/*
		Setup more routes, public routes should in routes/public
		Private routes that requires authentication of admin permission should go in routes/private
	*/
}
