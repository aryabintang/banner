package controllers

import (
	"context"
	"fmt"
	"golang_cms/configs"
	"golang_cms/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bannerCollection *mongo.Collection = configs.GetCollection(configs.DB, "Banner")
var validate = validator.New()

func CreateBanner(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var banner models.Banner
	defer cancel()

	//validate the request body
	if err := c.Bind(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  400,
			"Message": err.Error(),
		})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&banner); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  400,
			"Message": validationErr.Error(),
		})
		return
	}

	now := time.Now()
	banner.Id = now.UnixNano()
	result, err := bannerCollection.InsertOne(ctx, banner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  500,
			"Message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Data":    result,
		"Status":  200,
		"Message": "Data Berhasil Dibuat",
	})
}

func GetABanner(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	bannerId := c.Param("bannerId")
	var banner models.Banner
	defer cancel()
	fmt.Println(bannerId)
	err := bannerCollection.FindOne(ctx, bson.M{"id": bannerId[0]}).Decode(&banner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  500,
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": 200,
		"Data":   banner,
	})
}

func EditABanner(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	bannerId := c.Param("bannerId")
	var banner models.Banner
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(bannerId)

	//validate the request body
	if err = c.Bind(&banner); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  500,
			"Message": "Status Internal Server Error",
		})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&banner); validationErr != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  500,
			"Message": validationErr,
		})
		return

	}

	update := bson.M{"banner": banner.Banner, "alt": banner.Alt, "link": banner.Link}

	result, err := bannerCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  500,
			"Message": err,
		})
		return
	}

	//get updated Banner details
	var updatedBanner models.Banner
	if result.MatchedCount == 0 {
		err := bannerCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&banner)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Status":  500,
				"Message": "Internal Server Error",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{

		"Status":  200,
		"Message": "Data Berhasil Diubah",
		"Data":    updatedBanner,
	})

}

func DeleteBanner(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	bannerId := c.Param("bannerId")
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(bannerId)

	result, err := bannerCollection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  500,
			"Message": err,
		})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  404,
			"Message": result,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  200,
		"Message": "Data Berhasil Di Hapus",
	})

}

func GetAllBanner(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var banner []models.Banner
	defer cancel()

	results, err := bannerCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  500,
			"Message": err.Error(),
		})
		return
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleBanner models.Banner
		if err = results.Decode(&singleBanner); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Status":  500,
				"Message": err.Error(),
			})
			return
		}

		banner = append(banner, singleBanner)
	}

	c.JSON(http.StatusOK, gin.H{
		"Data":    banner,
		"Status":  200,
		"Message": "success",
	})

}
