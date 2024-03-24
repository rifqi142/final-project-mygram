package controller

import (
	"encoding/json"
	"errors"
	"final-project-mygram/config"
	"final-project-mygram/helper"
	"final-project-mygram/model"
	"final-project-mygram/util"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateComment(ctx *gin.Context) {
	db := config.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	commentRequest := model.CreateCommentRequest{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		// Bind JSON data to CreatePhotoRequest
		if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error(), "Invalid request"))
			return
		}
	} else {
		 // If content type is not JSON, respond with error
		 ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, "Invalid content type", "Content type must be application/json"))
		 return
	}

    // Check if photo_id exists in the database
    var photo model.Photo
    if err := db.First(&photo, commentRequest.PhotoID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, "Photo not found", "Photo ID does not exist"))
            return
        }
        ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to create comment"))
        return
    }
	
	comment := model.Comment{
		Message: commentRequest.Message,
		PhotoID: commentRequest.PhotoID,
		UserID: userID,
	}

	err := db.Debug().Create(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to create comment"))
		return
	}

	commentString, _ := json.Marshal(comment)
	commentResponse := model.CreateCommentResponse{}
	json.Unmarshal(commentString, &commentResponse)

	ctx.JSON(http.StatusCreated, util.CreateResponse(true, commentResponse, "", "Comment created successfully"))
}

func GetAllComment(ctx *gin.Context) {
	db := config.GetDB()

	comments := []model.Comment{}

	err := db.Debug().Preload("User").Preload("Photo").Order("id asc").Find(&comments).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to get comments"))
		return
	}

	commentsString, _ := json.Marshal(comments)
	commentsResponse := []model.CommentResponse{}
	json.Unmarshal(commentsString, &commentsResponse)

	ctx.JSON(http.StatusOK, util.CreateResponse(true, commentsResponse, "", "Comments fetched successfully"))
}

func GetCommentById(ctx *gin.Context) {
	db := config.GetDB()
	commentID := ctx.Param("CommentId")
	comment := model.Comment{}

	err := db.Debug().Preload("User").Preload("Photo").First(&comment, commentID).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, util.CreateResponse(false, nil, err.Error(), "Comment not found"))
		return
	}

	commentString, _ := json.Marshal(comment)
	commentResponse := model.CommentResponse{}
	json.Unmarshal(commentString, &commentResponse)

	ctx.JSON(http.StatusOK, util.CreateResponse(true, commentResponse, "", "Comment fetched successfully"))
}

func UpdateComment(ctx *gin.Context) {
	db := config.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	commentRequest := model.UpdateCommentRequest{}
	commentId, _ := strconv.Atoi(ctx.Param("CommentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		// Bind JSON data to CreatePhotoRequest
		if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error(), "Invalid request"))
			return
		}
	} else {
		 // If content type is not JSON, respond with error
		 ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, "Invalid content type", "Content type must be application/json"))
		 return
	}

	comment := model.Comment{}
	comment.ID = uint(commentId)
	comment.UserID = userID

	// Check if comment exists
	if err := db.Debug().First(&comment).Error; err != nil {
		ctx.JSON(http.StatusNotFound, util.CreateResponse(false, nil, err.Error(), "Comment not found"))
		return
	}

	updateString, _ := json.Marshal(commentRequest)
	updateData := model.Comment{}
	json.Unmarshal(updateString, &updateData)

	err := db.Debug().Model(&comment).Updates(updateData).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to update comment"))
		return
	}
	_ = db.First(&comment, comment.ID).Error

	// Marshal comment to JSON
	commentString, _ := json.Marshal(comment)
	commentResponse := model.UpdateCommentResponse{}
	json.Unmarshal(commentString, &commentResponse)

	ctx.JSON(http.StatusOK, util.CreateResponse(true, commentResponse, "", "Comment updated successfully"))
}

func DeleteComment(ctx *gin.Context) {
	db := config.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)

	commentID, err := strconv.Atoi(ctx.Param("CommentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error(), "Invalid comment ID"))
		return
	}

	userID := uint(userData["id"].(float64))

	comment := model.Comment{}
	if err := db.First(&comment, commentID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, util.CreateResponse(false, nil, err.Error(), "Comment not found"))
		return
	}

	if comment.UserID != userID {
		ctx.JSON(http.StatusForbidden, util.CreateResponse(false, nil, "Forbidden", "You are not authorized to perform this action"))
		return
	}

	if err := db.Delete(&comment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error(), "Failed to delete comment"))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, nil, "", "Your comment has been successfully deleted "))
}

