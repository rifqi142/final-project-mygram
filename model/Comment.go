package model

import (
	"time"
)

type Comment struct {
	ID uint `gorm:"primaryKey" json:"id"`
	UserID  uint `json:"user_id"`
	User *User `json:",omitempty"`
	PhotoID uint `json:"photo_id"`
	Photo *Photo `json:",omitempty"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//  Request for Comment
type CreateCommentRequest struct {
	Message string `json:"message" form:"message" valid:"required~Message is required" binding:"required"`
	PhotoID uint `json:"photo_id" form:"photo_id" valid:"required~Photo ID is required"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" form:"message" valid:"required~Message is required" binding:"required"`
}

// Custom Response Request for Comment
type CommentResponse struct {
	ID uint `json:"id"`
	Message string `json:"message"`
	PhotoID uint `json:"photo_id"`
	UserID uint `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	User *UserCommentResponse
	Photo *PhotoCommentResponse
}

type CreateCommentResponse struct {
	ID uint `json:"id"`
	Message string `json:"message"`
	PhotoID uint `json:"photo_id"`
	UserID uint `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateCommentResponse struct {
	ID uint `json:"id"`
	Message string `json:"message"`
	PhotoID uint `json:"photo_id"`
	UserID uint `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}