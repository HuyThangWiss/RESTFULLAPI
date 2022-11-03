package Config

import (
	"ProJectTest/BuildingAPI/ConnectApi"
	"ProJectTest/BuildingAPI/HashFunction"
	"ProJectTest/BuildingAPI/InformationAPI"
	"ProJectTest/BuildingAPI/getCollecttion"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Create(c *gin.Context)  {
	var DB=ConnectApi.ConnectDB()
	var postCollection = getCollecttion.GetCollection(DB,"Books")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	Books := new(InformationAPI.Books)
	defer cancel()

	if err := c.BindJSON(&Books); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	postLoad := InformationAPI.Books{
		Id:   Books.Id,
		Name: Books.Name,
		Year: Books.Year,
	}
	result,err := postCollection.InsertOne(ctx,postLoad)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated,gin.H{"Message ":"Post sucessfully","Data":map[string]interface{}{"Data ":result}})
}

func Find(c *gin.Context)  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB=ConnectApi.ConnectDB()
	var postCollection = getCollecttion.GetCollection(DB,"Books")
	Id := c.Param("Id")
	intvar, _ :=strconv.Atoi(Id)
	var result []InformationAPI.Books
	defer cancel()
	filter,err := postCollection.Find(ctx,bson.M{"Id":intvar})
	if err != nil{
		log.Fatal(err)
	}

	if err=filter.All(context.TODO(),&result);err != nil{
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated,gin.H{"Data  find":map[string]interface{}{"Data":result}})

}


func Select(c *gin.Context)  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB=ConnectApi.ConnectDB()
	var postCollection = getCollecttion.GetCollection(DB,"Books")
	var result []InformationAPI.Books
	defer cancel()
	curson2,err := postCollection.Find(context.TODO(),bson.M{})
	if err != nil{
		log.Fatal(err)
	}
	if err = curson2.All(ctx,&result);err != nil{
		log.Fatal(err)
	}

	c.JSON(http.StatusOK,gin.H{"Data ":map[string]interface{}{"Data":result}})
}

func Update(c *gin.Context)  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB=ConnectApi.ConnectDB()
	var postCollection = getCollecttion.GetCollection(DB,"Books")
	var result InformationAPI.Books
	defer cancel()
	Id := c.Param("Id")
	intvar, _ :=strconv.Atoi(Id)

	if err := c.BindJSON(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	edited := bson.M{"Name":result.Name,"Year":result.Year}
	books,err := postCollection.UpdateOne(ctx,bson.M{"Id":intvar},bson.M{"$set":edited})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	if books.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"Data ":books.ModifiedCount})
}


func Delete(c *gin.Context)  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = ConnectApi.ConnectDB()
	Id := c.Param("Id")
	intvar, _ :=strconv.Atoi(Id)
	var postCollection = getCollecttion.GetCollection(DB, "Books")
	defer cancel()
	var result InformationAPI.Books

	if err := c.BindJSON(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	filter := bson.D{{"Id", bson.D{{"$", intvar}}}}
		books, err := postCollection.DeleteOne(ctx, filter)
		if err != nil{
			log.Fatal(err)
		}
	c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully", "Data": books.DeletedCount})
}



func CreateHash(c *gin.Context)  {
	var DB=ConnectApi.ConnectDB()
	var postCollection = getCollecttion.GetCollection(DB,"Books")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	Books := new(InformationAPI.Books)
	defer cancel()

	if err := c.BindJSON(&Books); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	hashedPassword, _ := HashFunction.HashPassword(Books.Name)
	Books.Name = hashedPassword
	postLoad := InformationAPI.Books{
		Id:   Books.Id,
		Name: Books.Name,
		Year: Books.Year,
	}
	result,err := postCollection.InsertOne(ctx,postLoad)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated,gin.H{"Message ":"Post sucessfully","Data":map[string]interface{}{"Data ":result}})
}


















