package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"title" valid:"required~Message is required"`
	UserID  uint   `gorm:"not null" json:"user_id" valid:"required~userID is required"`
	PhotoID uint   `gorm:"not null" json:"photo_id" valid:"required~photoID is required"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	err = errCreate
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(c)

	err = errUpdate
	return
}
