package models

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code" validate:"required,min=2,max=3"`
}
