package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name             string `gorm:"not null" json:"name" valid:"required~Your name is required"`
	Social_Media_Url string `gorm:"not null;uniqueIndex" json:"social_media_url" form:"social_media_form" valid:"required~Your social media URL is required"`
	UserID           uint   `gorm:"not null" json:"user_id" valid:"required~UserID is required"`
	User             *User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	err = errCreate
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(s)

	err = errUpdate
	return
}
