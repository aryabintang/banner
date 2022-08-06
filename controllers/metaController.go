package controllers

import (
	"context"
	"fmt"
	"golang_cms/configs"
	"golang_cms/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var metaCollection *mongo.Collection = configs.GetCollection(configs.DB, "Meta")
var val = validator.New()

func Createmeta(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var meta models.Meta
	defer cancel()

	//val the request body
	if err := c.Bind(&meta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  400,
			"Message": err.Error(),
		})
		return
	}

	//use the validator library to val required fields
	if validationErr := val.Struct(&meta); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  400,
			"Message": validationErr.Error(),
		})
		return
	}

	now := time.Now()
	meta.Id = int(now.UnixNano())
	result, err := metaCollection.InsertOne(ctx, meta)
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

func GetAmeta(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	metaId := c.Param("metaId")
	var meta models.Meta
	defer cancel()

	fmt.Println(metaId)
	i, _ := strconv.Atoi(metaId)

	err := metaCollection.FindOne(ctx, bson.M{"id": i}).Decode(&meta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  500,
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{

		"Data": meta,
	})
}

func EditAmeta(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	metaId := c.Param("metaId")
	var meta models.Meta
	defer cancel()

	fmt.Println(metaId)
	i, _ := strconv.Atoi(metaId)

	//val the request body
	if err := c.Bind(&meta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	//use the validator library to val required fields
	if validationErr := val.Struct(&meta); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  500,
			"Message": validationErr.Error(),
		})
	}

	update := bson.M{"Title": meta.Title, "descripsi": meta.Descrpsi, "Kategori Produk": meta.Kategori_produk}

	result, err := metaCollection.UpdateOne(ctx, bson.M{"id": i}, bson.M{"$set": update})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  400,
			"Message": err.Error(),
		})
	}
	//get updated user details
	var updatedUser models.Meta
	if result.MatchedCount == 1 {
		err := metaCollection.FindOne(ctx, bson.M{"id": i}).Decode(&updatedUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Status":  400,
				"Message": err.Error(),
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "success",
		"Data":    updatedUser,
	})
}

func Deletemeta(c gin.Context) gin.HandlerFunc {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	metaId := c.Param("metaId")
	defer cancel()
	fmt.Println(metaId)
	i, _ := strconv.Atoi(metaId)

	result, err := metaCollection.DeleteOne(ctx, bson.M{"id": i})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  500,
			"Message": err,
		})

	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  404,
			"Message": result,
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  200,
		"Message": "Data Berhasil Di Hapus",
	})

	return Deletemeta(gin.Context{})
}

func GetAllmeta(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var meta []models.Meta
	defer cancel()

	results, err := metaCollection.Find(ctx, bson.M{})

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
		var singlemeta models.Meta
		if err = results.Decode(&singlemeta); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Status":  500,
				"Message": err.Error(),
			})
			return
		}

		meta = append(meta, singlemeta)
	}

	c.JSON(http.StatusOK, gin.H{
		"Data":    meta,
		"Status":  200,
		"Message": "success",
	})

}
