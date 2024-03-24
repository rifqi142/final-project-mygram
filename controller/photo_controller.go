// GetAll, GetOne, createPhoto, UpdatePhoto, DeletePhoto

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


func CreatePhoto(ctx *gin.Context) {
	db := config.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	photoRequest := model.CreatePhotoRequest{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		// Bind JSON data to CreatePhotoRequest
		if err := ctx.ShouldBindJSON(&photoRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error(), "Invalid request"))
			return
		}
	} else {
		 // If content type is not JSON, respond with error
		 ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, "Invalid content type", "Content type must be application/json"))
		 return
	}

	photo := model.Photo{
		Title: photoRequest.Title,
		Caption: photoRequest.Caption,
		PhotoURL: photoRequest.PhotoURL,
		UserID: userID,
	}

	err := db.Debug().Create(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to create photo"))
		return
	}
	_ = db.First(&photo, photo.ID).Error

	photoString, _ := json.Marshal(photo)
	photoResponse := model.CreatePhotoResponse{}
	json.Unmarshal(photoString, &photoResponse)

	ctx.JSON(http.StatusCreated, util.CreateResponse(true, photoResponse, "", "Photo created successfully"))
}

func GetAllPhoto(ctx *gin.Context) {
	db := config.GetDB()
	photos := []model.Photo{}

	err := db.Debug().Preload("User").Find(&photos).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to fetch photos"))
		return
	}

	photoString, _ := json.Marshal(photos)
	photoResponse := []model.GetPhotoResponse{}
	json.Unmarshal(photoString, &photoResponse)

	ctx.JSON(http.StatusOK, util.CreateResponse(true, photoResponse, "", "Photos fetched successfully"))
}

func GetOnePhoto(ctx *gin.Context) {
	db := config.GetDB()
	photoID := ctx.Param("PhotoId")
	photo := model.Photo{}

	err := db.Debug().Preload("User").First(&photo, photoID).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, util.CreateResponse(false, nil, err.Error(), "Photo not found"))
		return
	}

	photoString, _ := json.Marshal(photo)
	photoResponse := model.GetPhotoResponse{}
	json.Unmarshal(photoString, &photoResponse)

	ctx.JSON(http.StatusOK, util.CreateResponse(true, photoResponse, "", "Photo fetched successfully"))
}

func UpdatePhoto(ctx *gin.Context) {
	db := config.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	photoRequest := model.UpdatePhotoRequest{}
	photoId, _ := strconv.Atoi(ctx.Param("PhotoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		// Bind JSON data to UpdatePhotoRequest
		if err := ctx.ShouldBindJSON(&photoRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error(), "Invalid request"))
			return
		}
	} else {
		 // If content type is not JSON, respond with error
		 ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, "Invalid content type", "Content type must be application/json"))
		 return
	}

	photo := model.Photo{}
	photo.ID = uint(photoId)
	photo.UserID = userID

	// Check if photo exists
	if err := db.First(&photo, photo.ID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, util.CreateResponse(false, nil, err.Error(), "Photo not found"))
		return
	}

	// Marshal photoRequest to string
	updateString, _ := json.Marshal(photoRequest)
	updateData := model.Photo{}
	json.Unmarshal(updateString, &updateData)

	// Perform update
	if err := db.Model(&photo).Updates(updateData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to update photo"))
		return
	}

	// Fetch updated photo
	updatedPhoto := model.Photo{}
	if err := db.First(&updatedPhoto, photo.ID).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to fetch updated photo"))
		return
	}

	photoString, _ := json.Marshal(updatedPhoto)
	photoResponse := model.UpdatePhotoResponse{}
	json.Unmarshal(photoString, &photoResponse)

	ctx.JSON(http.StatusOK, util.CreateResponse(true, photoResponse, "", "Photo updated successfully"))
}

func DeletePhoto(ctx *gin.Context) {
	db := config.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)

	photoId, _ := strconv.Atoi(ctx.Param("PhotoId"))
	userID := uint(userData["id"].(float64))

	photo := model.Photo{}
	photo.ID = uint(photoId)
	photo.UserID = userID

	err := db.Delete(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to delete photo"))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, nil, "", "Photo deleted successfully"))
}