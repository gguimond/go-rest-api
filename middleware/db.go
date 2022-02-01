package middleware

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB(database *mongo.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", database)
		ctx.Next()
	}
}
