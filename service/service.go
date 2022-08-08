package service

import (
	"context"
	"fmt"
	"golang_cms/configs"
	"golang_cms/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var collection *mongo.Collection = configs.GetCollection(configs.DB, "user")

//Userservice is to handle user relation db query
type Userservice struct{}

//Create is to register new user
func (userservice Userservice) Create(user *(models.User), c gin.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userEmail := c.Param("userEmail")
	var data models.User
	defer cancel()

	fmt.Println(userEmail)

	err := collection.FindOne(ctx, bson.M{"Email": user.Email}).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  500,
			"Message": err.Error(),
		})

	}

	c.JSON(http.StatusOK, gin.H{

		"Data": data,
	})

	return err
}
