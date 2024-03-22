package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username"`
	Email   string `gorm:"not null;uniqueIndex" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
	Age int `gorm:"not null" json:"age" form:"age"`
	Photos []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

