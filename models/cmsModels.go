package models

import (
	"golang_cms/configs"

	"github.com/golang-jwt/jwt"
	"github.com/goonode/mogo"
)

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
	Id       int    `json:"id,omitempty"`
	Fname    string `json:"fname,omitempty" validasi:"required"`
	Lname    string `json:"lname,omitempty" validasi:"required"`
	Email    string `json:"email,omitempty" validasi:"required"`
	Password string `json:"password,omitempty" validasi:"required"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqBody struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required,gte=6"`
}

type ResBody struct {
	Token string `json:"token"`
}

type Claims struct {
	Email string `json:"email"`
	*jwt.StandardClaims
}

//GetJwtToken returns jwt token with user email claims
func (user *User) GetJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": string(user.Email),
	})
	secretKey := configs.EnvMongoURI("TOKEN_KEY", "")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

func init() {
	mogo.ModelRegistry.Register(User{})
}
