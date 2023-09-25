package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string `gorm:"not null" json:"title" valid:"required~Title is required"`
	Caption   string `gorm:"not null" json:"caption" valid:"required~Caption is required"`
	Photo_Url string `gorm:"not null" json:"photo_url" valid:"required~Photo URL is required"`
	UserID    uint   `gorm:"not null" json:"user_id" valid:"required~userID is required"`
	Comments  []Comment
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	err = errCreate
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	err = errUpdate
	return
}
