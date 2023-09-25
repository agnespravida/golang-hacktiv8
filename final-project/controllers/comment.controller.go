package controllers

import (
	"final-project/databases"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CommentCreate(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)
	user := c.MustGet("user").(jwt.MapClaims)

	Comment := models.Comment{}
	UserID := uint(user["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	Comment.UserID = UserID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)

}

func CommentGetAll(c *gin.Context) {
	db := databases.GetDB()
	Comment := []models.Comment{}

	err := db.Debug().Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)

}

func CommentGetByUserLogin(c *gin.Context) {
	db := databases.GetDB()
	Comment := models.Comment{}
	user := c.MustGet("user").(jwt.MapClaims)
	UserID := uint(user["id"].(float64))

	err := db.Debug().Where("user_id = ?", UserID).Take(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)

}
func CommentGetByID(c *gin.Context) {
	db := databases.GetDB()
	Comment := models.Comment{}
	CommentID, err := strconv.Atoi(c.Param("commentID"))

	err = db.Debug().Where("id = ?", CommentID).Take(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)

}

func UpdateComment(c *gin.Context) {
	db := databases.GetDB()
	user := c.MustGet("user").(jwt.MapClaims)
	UserID := uint(user["id"].(float64))
	CommentID, err := strconv.Atoi(c.Param("commentID"))
	Comment := models.Comment{}
	contentType := helpers.GetContentType(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   "Invalid Parameter",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = UserID
	Comment.ID = uint(CommentID)

	err = db.Model(&Comment).Where("id = ?", CommentID).Updates(models.Comment{PhotoID: Comment.PhotoID, Message: Comment.Message}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

func DeleteComment(c *gin.Context) {
	db := databases.GetDB()
	CommentID, err := strconv.Atoi(c.Param("commentID"))
	Comment := models.Comment{}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   "Invalid Parameter",
		})
		return
	}

	err = db.Model(&Comment).Where("id = ?", CommentID).Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Success Delete Data",
	})
}
