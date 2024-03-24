package model

import "time"

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url"`
	UserID         uint   ` json:"user_id"`
	User           *User
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

// Request for SocialMedia
type CreateSocialMediaRequest struct {
	Name           string `json:"name" form:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" form:"social_media_url" binding:"required"`
}


type UpdateSocialMediaRequest struct {
	Name           string `json:"name" form:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" form:"social_media_url" binding:"required"`
}

// Custom Response Request for SocialMedia
type GetSocialMediaResponse struct {
	ID 		  	   uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaURL string     `json:"social_media_url"`
	UserID         uint   	  `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt 	   *time.Time `json:"updated_at"`
	User           *UserSocialMediaResponse
}

type CreateSocialMediaResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaURL string     `json:"social_media_url"`
	UserID         uint   	  `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
}

type UpdateSocialMediaResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaURL string     `json:"social_media_url"`
	UserID         uint   	  `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
