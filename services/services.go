package services

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"newproject/data"
	"newproject/database"
)

func GetDetail(c *gin.Context) {
	var info data.Info
	head := c.Request.Header
	name := head.Get("name")
	client := database.GetClient()
	collection := client.Database("newproject").Collection("data")
	result := collection.FindOne(c, bson.M{"name": name})
	_ = result.Decode(&info)
	c.JSON(http.StatusOK, info)
}
func GetDetails(c *gin.Context) {
	var infos []data.Info
	client := database.GetClient()
	collection := client.Database("newproject").Collection("data")
	result, _ := collection.Find(c, bson.M{})
	for result.Next(c) {
		var info data.Info
		_ = result.Decode(&info)
		infos = append(infos, info)
	}
	c.JSON(http.StatusOK, infos)

	
}
func PostDetails(c *gin.Context) {
	var info data.Info
	_ = c.BindJSON(&info)
	client := database.GetClient()
	collection := client.Database("newproject").Collection("data")
	id, _ := collection.InsertOne(c, info)
	c.JSON(http.StatusOK, gin.H{
		"created successfully ": id,
	})
}
func PutDetails(c *gin.Context) {
	var info data.Info
	head := c.Request.Header
	name := head.Get("name")
	_ = c.BindJSON(&info)
	client := database.GetClient()
	collection := client.Database("newproject").Collection("data")
	id, _ := collection.UpdateOne(c, bson.M{"name": name}, bson.M{"$set": info})
	c.JSON(http.StatusOK, gin.H{
		"updated ": id,
	})
}
func DeleteDetails(c *gin.Context) {
	var info data.Info
	client := database.GetClient()
	_ = c.BindJSON(&info)
	collection := client.Database("newproject").Collection("data")
	res, _ := collection.DeleteOne(c, bson.M{"name": info.Name})
	c.JSON(http.StatusOK, gin.H{
		"Deleted ": res,
	})
}
