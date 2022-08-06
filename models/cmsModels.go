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

type User struct {
	Email    string `json:"email,,omitempty" validasi:"required"`
	Password string `json:"password,,omitempty" validasi:"required"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
