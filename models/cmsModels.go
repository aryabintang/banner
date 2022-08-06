package models

type Banner struct {
	Id     int    `json:"id,omitempty"`
	Banner string `json:"banner,omitempty" validate:"required"`
	Alt    string `json:"alt,omitempty" validate:"required"`
	Link   string `json:"link,omitempty" validate:"required"`
}

type Meta struct {
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty" valid:"required"`
	Descrpsi string `json:"descrpsi,omitempty" val:"required"`
	Kategori string `json:"kategori,omitempty" val:"required"`
}
