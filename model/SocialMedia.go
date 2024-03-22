package model

import "gorm.io/gorm"

type SocialMedia struct {
	gorm.Model
	Name string `gorm:"not null" json:"name" form:"name"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url"`
	UserID uint ` json:"user_id"`
	User *User
}