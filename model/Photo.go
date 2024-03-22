package model

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption string `gorm:"not null" form:"caption" valid:"-"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	UserID uint `json:"user_id"`
	User *User `json:",omitempty"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}