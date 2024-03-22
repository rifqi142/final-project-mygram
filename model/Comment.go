package model

// import (
// 	"gorm.io/gorm"
// )

// type Comment struct {
// 	gorm.Model
// 	UserID  uint `json:"user_id"`
// 	User *User `json:",omitempty"`
// 	PhotoID uint `json:"photo_id"`
// 	Photo *Photo `json:",omitempty"`
// 	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
// }