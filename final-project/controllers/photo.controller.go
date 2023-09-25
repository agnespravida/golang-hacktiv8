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

func PhotoCreate(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)
	user := c.MustGet("user").(jwt.MapClaims)

	Photo := models.Photo{}
	UserID := uint(user["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo.UserID = UserID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)

}

func PhotoGetAll(c *gin.Context) {
	db := databases.GetDB()
	Photo := []models.Photo{}

	err := db.Debug().Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)

}

func PhotoGetByUserLogin(c *gin.Context) {
	db := databases.GetDB()
	Photo := models.Photo{}
	user := c.MustGet("user").(jwt.MapClaims)
	UserID := uint(user["id"].(float64))

	err := db.Debug().Where("user_id = ?", UserID).Take(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)

}
func PhotoGetByID(c *gin.Context) {
	db := databases.GetDB()
	Photo := models.Photo{}
	photoID, err := strconv.Atoi(c.Param("photoID"))

	err = db.Debug().Where("id = ?", photoID).Take(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)

}

func UpdatePhoto(c *gin.Context) {
	db := databases.GetDB()
	user := c.MustGet("user").(jwt.MapClaims)
	UserID := uint(user["id"].(float64))
	photoID, err := strconv.Atoi(c.Param("photoID"))
	Photo := models.Photo{}
	contentType := helpers.GetContentType(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   "Invalid Parameter",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = UserID
	Photo.ID = uint(photoID)

	err = db.Model(&Photo).Where("id = ?", photoID).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, Photo_Url: Photo.Photo_Url}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func DeletePhoto(c *gin.Context) {
	db := databases.GetDB()
	photoID, err := strconv.Atoi(c.Param("photoID"))
	Photo := models.Photo{}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   "Invalid Parameter",
		})
		return
	}

	err = db.Model(&Photo).Where("id = ?", photoID).Delete(&Photo).Error
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
