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

func SocialMediaCreate(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)
	user := c.MustGet("user").(jwt.MapClaims)

	SocialMedia := models.SocialMedia{}
	UserID := uint(user["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}
	SocialMedia.UserID = UserID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)

}

func SocialMediaGetAll(c *gin.Context) {
	db := databases.GetDB()
	SocialMedia := []models.SocialMedia{}

	err := db.Debug().Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)

}

func SocialMediaGetByUserLogin(c *gin.Context) {
	db := databases.GetDB()
	SocialMedia := models.SocialMedia{}
	user := c.MustGet("user").(jwt.MapClaims)
	UserID := uint(user["id"].(float64))

	err := db.Debug().Where("user_id = ?", UserID).Take(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)

}
func SocialMediaGetByID(c *gin.Context) {
	db := databases.GetDB()
	SocialMedia := models.SocialMedia{}
	socmedID, err := strconv.Atoi(c.Param("socmedID"))

	err = db.Debug().Where("id = ?", socmedID).Take(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)

}

func UpdateSocialMedia(c *gin.Context) {
	db := databases.GetDB()
	user := c.MustGet("user").(jwt.MapClaims)
	UserID := uint(user["id"].(float64))
	socmedID, err := strconv.Atoi(c.Param("socmedID"))
	SocialMedia := models.SocialMedia{}
	contentType := helpers.GetContentType(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   "Invalid Parameter",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = UserID
	SocialMedia.ID = uint(socmedID)

	err = db.Model(&SocialMedia).Where("id = ?", socmedID).Updates(models.SocialMedia{Name: SocialMedia.Name, Social_Media_Url: SocialMedia.Social_Media_Url}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	db := databases.GetDB()
	socmedID, err := strconv.Atoi(c.Param("socmedID"))
	SocialMedia := models.SocialMedia{}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   "Invalid Parameter",
		})
		return
	}

	err = db.Model(&SocialMedia).Where("id = ?", socmedID).Delete(&SocialMedia).Error
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
