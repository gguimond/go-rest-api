package controllers

import (
	"go-rest-api/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetEpisodes(ctx *gin.Context) {
	database := ctx.MustGet("db").(*mongo.Database)
	episodesCollection := database.Collection("episodes")

	var episodes []types.Episode
	cursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &episodes); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": episodes})
}

type InsertEpisodeInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Duration    int32  `json:"duration" binding:"required"`
}

func InsertOneEpisode(ctx *gin.Context) {
	var input InsertEpisodeInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := ctx.MustGet("db").(*mongo.Database)
	episodesCollection := database.Collection("episodes")

	episode := types.Episode{Title: input.Title, Description: input.Description, Duration: input.Duration}

	insertResult, err := episodesCollection.InsertOne(ctx, episode)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": insertResult})
}

func UpdateOneEpisode(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	var input InsertEpisodeInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := ctx.MustGet("db").(*mongo.Database)
	episodesCollection := database.Collection("episodes")

	insertResult, err := episodesCollection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.D{{"$set", input}})
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": insertResult})
}

func DeleteOneEpisode(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	database := ctx.MustGet("db").(*mongo.Database)
	episodesCollection := database.Collection("episodes")

	_, err = episodesCollection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
