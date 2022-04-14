# cards-api

A sample project for job offer on Toggl

## Introduction

This projects uses:
- github.com/go-playground/validator for validating requests
- github.com/gofiber/fiber for listening and responding the requestes
- github.com/google/uuid to create a unique deckId
- github.com/joho/godotenv to load simple configuration from root/.env

### Compiling and running

- Go to the root of the project and in the terminal run the command `go mod tidy`
- Configure the .env file, if you intend to test locally you should leave `ADDRESS=localhost` but if you intend to test in a server change localhost to the public ipv4 address of the server and make sure it has permission to listen to the specified port
- That's all if go is propperly installed just run `go run main.go` from the root directory

### Postman Requests for testing

The easiest way of testing is using postman, use the link  [Postman collection](https://www.postman.com/crimson-flare-4229/workspace/cards-api/collection/10761195-6f9ba5ca-66c5-4a0f-b006-4bbff1f5dab2?action=share&creator=10761195)

#### Create New Deck
Address: `localhost:9000/create-new-deck`

A post request to create a new deck and add it to the memory, body parameters:
```
{
    "cards": "AS,2S,3S",
    "shuffled": true
}
```
- cards should be cardCodes separated by commas or an empty string to create a full size deck
- shuffled is a boolean to indicate if the deck should or should not be shuffled

#### Draw Cards
Address: `localhost:9000/draw-cards`

A put request to draw cards from a specified deck, body parameters: 
```
{
    "id": "b530fe1f-bc3b-495c-b696-849082331818",
    "number": 2
}
```
- id should be the id returned by the response when calling Create New Deck
- number should be the number of cards to be drawn, if deck has insufficient cards it will return an error response.

#### Open Deck
Address: `localhost:9000/open-deck/{{deckId}}`
A Get request should return a response with deck info or error if deckId is not found

#### Get all decks (EXTRA)
Address: `localhost:9000/get-all-decks`

For debugging purposes there's a method to list all decks

## Remote Debugging

Feel free to my personal EC instance for testing request, Follow the link to the postman collection pointing to it

[Postman collection]([Postman collection](https://www.postman.com/crimson-flare-4229/workspace/cards-api/collection/10761195-6f9ba5ca-66c5-4a0f-b006-4bbff1f5dab2?action=share&creator=10761195))

