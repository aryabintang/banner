package models

type Banner struct {
	Id     int    `json:"id,omitempty"`
	Banner string `json:"banner,omitempty" validate:"required"`
	Alt    string `json:"alt,omitempty" validate:"required"`
	Link   string `json:"link,omitempty" validate:"required"`
}

type Meta struct {
	Id              int    `json:"id,omitempty"`
	Title           string `json:"tile,omitempty" validate:"required"`
	Descrpsi        string `json:"descrpsi,omitempty" validate:"required"`
	Kategori_produk string `json:"kategori_produk,omitempty" validate:"required"`
}
