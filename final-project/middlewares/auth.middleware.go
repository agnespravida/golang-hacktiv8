package middlewares

import (
	"final-project/databases"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("user", verifyToken)
		c.Next()
	}
}

func SocialMediaAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := databases.GetDB()
		socmedID, err := strconv.Atoi(c.Param("socmedID"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
				"error":   "Invalid Parameter",
			})
			return
		}
		user := c.MustGet("user").(jwt.MapClaims)
		userID := uint(user["id"].(float64))
		SocialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, uint(socmedID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Not Found",
				"error":   "Data not found",
			})
			return
		}
		if SocialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "You cannot access this data",
			})
			return
		}
		c.Next()
	}
}

func PhotoAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := databases.GetDB()
		photoID, err := strconv.Atoi(c.Param("photoID"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
				"error":   "Invalid Parameter",
			})
			return
		}
		user := c.MustGet("user").(jwt.MapClaims)
		userID := uint(user["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Not Found",
				"error":   "Data not found",
			})
			return
		}
		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "You cannot access this data",
			})
			return
		}
		c.Next()
	}
}

func CommentAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := databases.GetDB()
		commentID, err := strconv.Atoi(c.Param("commentID"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
				"error":   "Invalid Parameter",
			})
			return
		}
		user := c.MustGet("user").(jwt.MapClaims)
		userID := uint(user["id"].(float64))
		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Not Found",
				"error":   "Data not found",
			})
			return
		}
		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "You cannot access this data",
			})
			return
		}
		c.Next()
	}
}
