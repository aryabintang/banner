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
	Descrpsi string `json:"descrpsi,omitempty" valid:"required"`
	Kategori string `json:"kategori,omitempty" valid:"required"`
}

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}
