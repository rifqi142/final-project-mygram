package middleware

import (
	"final-project-mygram/helper"
	"final-project-mygram/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func (c *gin.Context) {
		verifyToken, err := helper.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.CreateResponse(false, nil, err.Error(), "Unauthenticated"))
			return
		}
		c.Set("userInfo", verifyToken)
		c.Next()
	}
}

// func PhotoAuthorization() gin.HandlerFunc {
// 	return func (c *gin.Context) {
// 		db := config.GetDB()
// 		photoId, err := strconv.Atoi(c.Param("PhotoId"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error(), "Invalid Photo ID"))
// 			return
// 		}

// 		userData := c.MustGet("userInfo").(jwt.MapClaims)
// 		userID := uint(userData["id"].(float64))
// 		photo := model.Photo{}

// 		err = db.Select("user_id").First(&photo, uint(photoId)).Error
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, util.CreateResponse(false, nil, err.Error(), "Photo not found"))
// 			return
// 		}

// 		if photo.UserID != userID {
// 			c.JSON(http.StatusForbidden, util.CreateResponse(false, nil, "Forbidden", "You are not authorized to perform this action"))
// 			return
// 		}

// 		c.Next()
// 	}
// }

// func CommentAuthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := config.GetDB()
// 		commentID, err := strconv.Atoi(c.Param("CommentId"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error(), "Invalid Comment ID"))
// 			return
// 		}
		
// 		userData := c.MustGet("userInfo").(jwt.MapClaims)
// 		userID := uint(userData["id"].(float64))
		
// 		comment := model.Comment{}

// 		err = db.Select("user_id").First(&comment, uint(commentID)).Error
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, util.CreateResponse(false, nil, err.Error(), "Comment not found"))
// 			return
// 		}

// 		// Memeriksa apakah pengguna memiliki otorisasi untuk menghapus komentar
// 		if comment.UserID != userID {
// 			c.JSON(http.StatusForbidden, util.CreateResponse(false, nil, "Forbidden", "You are not authorized to perform this action"))
// 			return
// 		}

// 		// Jika pengguna memiliki otorisasi, lanjutkan ke handler berikutnya
// 		c.Next()
// 	}
// }