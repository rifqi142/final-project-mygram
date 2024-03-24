package controller

import (
	"encoding/json"
	"final-project-mygram/config"
	"final-project-mygram/helper"
	"final-project-mygram/model"
	"final-project-mygram/util"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(ctx *gin.Context) {
	db := config.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	socialMediaRequest := model.CreateSocialMediaRequest{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		// Bind JSON data to CreatePhotoRequest
		if err := ctx.ShouldBindJSON(&socialMediaRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error(), "Invalid request"))
			return
		}
	} else {
		 // If content type is not JSON, respond with error
		 ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, "Invalid content type", "Content type must be application/json"))
		 return
	}

	socialMedia := model.SocialMedia{
		Name: socialMediaRequest.Name,
		SocialMediaURL: socialMediaRequest.SocialMediaURL,
		UserID: userID,
	}

	err := db.Debug().Create(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to create social media"))
		return
	}

	socialMediaString, _ := json.Marshal(socialMedia)
	socialMediaResponse := model.CreateSocialMediaResponse{}
	json.Unmarshal(socialMediaString, &socialMediaResponse)

	ctx.JSON(http.StatusOK, util.CreateResponse(true, socialMediaResponse, "Social media created successfully", ""))
}

func GetAllSocialMedia(ctx *gin.Context) {
	db := config.GetDB()

	socialMedia := []model.SocialMedia{}

	err := db.Debug().Preload("User").Order("id asc").Find(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to get social media"))
		return
	}

	socialMediaString, _ := json.Marshal(socialMedia)
	socialMediaResponse := []model.GetSocialMediaResponse{}
	json.Unmarshal(socialMediaString, &socialMediaResponse)

	ctx.JSON(http.StatusOK, util.CreateResponse(true, socialMediaResponse, "", "Social media retrieved successfully"))
}


func GetSocialMediaById(ctx *gin.Context) {
	db := config.GetDB()
	socialMediaID := ctx.Param("SocialMediaId")
	socialMedia := []model.SocialMedia{}

	err := db.Debug().Preload("User").First(&socialMedia, socialMediaID).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to get social media"))
		return
	}

	socialMediaString, _ := json.Marshal(socialMedia)
	socialMediaResponse := []model.GetSocialMediaResponse{}
	json.Unmarshal(socialMediaString, &socialMediaResponse)

	ctx.JSON(http.StatusOK, util.CreateResponse(true, socialMediaResponse, "", "Social media retrieved successfully"))
}

func UpdateSocialMedia(ctx *gin.Context) {
	db := config.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	socialMediaRequest := model.UpdateSocialMediaRequest{}
	socialMediaId, _ := strconv.Atoi(ctx.Param("SocialMediaId"))
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		// Bind JSON data to CreatePhotoRequest
		if err := ctx.ShouldBindJSON(&socialMediaRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error(), "Invalid request"))
			return
		}
	} else {
		 // If content type is not JSON, respond with error
		 ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, "Invalid content type", "Content type must be application/json"))
		 return
	}

	socialMedia := model.SocialMedia{}
	socialMedia.ID = uint(socialMediaId)
	socialMedia.UserID = UserID

	// Check if social media exists
	err := db.Debug().First(&socialMedia, socialMedia.ID).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, util.CreateResponse(false, nil, err.Error(), "Social media not found"))
		return
	}

	if socialMedia.UserID != UserID {
		ctx.JSON(http.StatusUnauthorized, util.CreateResponse(false, nil, "Unauthorized", "You are not authorized to update this social media"))
		return
	}

	updateString, _ := json.Marshal(socialMediaRequest)
	updateData := model.SocialMedia{}
	json.Unmarshal(updateString, &updateData)

	err = db.Debug().Model(&socialMedia).Updates(updateData).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to update social media"))
		return
	}
	_ = db.First(&socialMedia, socialMedia.ID).Error

	socialMediaString, _ := json.Marshal(socialMedia)
	socialMediaResponse := model.UpdateSocialMediaResponse{}
	json.Unmarshal(socialMediaString, &socialMediaResponse)

	ctx.JSON(http.StatusOK, util.CreateResponse(true, socialMediaResponse, "", "Social media updated successfully"))
}

func DeleteSocialMedia(ctx *gin.Context) {
	db := config.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)

	socialMediaID, err := strconv.Atoi(ctx.Param("SocialMediaId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error(), "Invalid social media ID"))
		return
	}

	userID := uint(userData["id"].(float64))

	socialMedia := model.SocialMedia{}
	if err := db.First(&socialMedia, socialMediaID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, util.CreateResponse(false, nil, err.Error(), "Social media not found"))
		return
	}

	if socialMedia.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, util.CreateResponse(false, nil, "Unauthorized", "You are not authorized to delete this social media"))
		return
	}

	if err := db.Delete(&socialMedia).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to delete social media"))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, nil, "", "your social media has been successfully deleted"))
}