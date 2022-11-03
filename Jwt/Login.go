package Jwt

import (
	"ProJectTest/BuildingAPI/Auth"
	"ProJectTest/BuildingAPI/ConnectApi"
	"ProJectTest/BuildingAPI/InformationAPI"
	"ProJectTest/BuildingAPI/getCollecttion"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func Login(c *gin.Context)  {
	var DB=ConnectApi.ConnectDB()
	var postCollection = getCollecttion.GetCollection(DB,"Books")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	var Input InformationAPI.Books
	var Admin InformationAPI.Books
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	var err = postCollection.FindOne(ctx,bson.M{
		"Id":Input.Id,
		"Name":Input.Name,
	}).Decode(&Admin)
	if err != nil{
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK,gin.H{"Login ":"Successfully"})
}

type TokenRequest struct {
	Id   int64  `json:"Id" bson:"Id" binding:"required`
	Name string `json:"Name" bson:"Name" binding:"required `
}

func LoginHash(c *gin.Context)  {
	var DB=ConnectApi.ConnectDB()
	var postCollection = getCollecttion.GetCollection(DB,"Books")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	var Input InformationAPI.Books
	var Admin InformationAPI.Books
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	var err = postCollection.FindOne(ctx,bson.M{
		"Id":Input.Id,
	}).Decode(&Admin)
	if err != nil{
		log.Fatal(err)
		return
	}

	err2 := bcrypt.CompareHashAndPassword([]byte(Admin.Name), []byte(Input.Name))
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{"Login ": "Fail"})
		return
	}
	token,err3 := Auth.GenerateJWT(Admin.Id,Admin.Name)
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"Token ":token})
}























