package model

import (
	"time"
)

type Photo struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption string `gorm:"not null" form:"caption" valid:"required~Caption is required" json:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	UserID uint `json:"user_id"`
	User *User `json:",omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

//  Request for Photo
type CreatePhotoRequest struct {
	Title string `json:"title" from:"title" binding:"required"`
	Caption string `json:"caption" from:"caption" binding:"required"`
	PhotoURL string `json:"photo_url" from:"photo_url" binding:"required"`
}

type UpdatePhotoRequest struct  {
	Title string `json:"title" from:"title" binding:"required"`
	Caption string `json:"caption" from:"caption"`
	PhotoURL string `json:"photo_url" from:"photo_url" binding:"required"`
}

// Custom Response Request for Photo
type GetPhotoResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID uint `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User *UserPhotoResponse
}

type CreatePhotoResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID uint `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type UpdatePhotoResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID uint `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type PhotoCommentResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID uint `json:"user_id"`
}