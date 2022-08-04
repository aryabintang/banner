package models

type Banner struct {
	Id     int    `json:"id,omitempty"`
	Banner string `json:"banner,omitempty" validate:"required"`
	Alt    string `json:"alt,omitempty" validate:"required"`
	Link   string `json:"link,omitempty" validate:"required"`
}
