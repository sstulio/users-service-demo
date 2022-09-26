package app

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sstulio/users-service-demo/internal/models"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User

		err := db.Model(&models.User{}).Find(&users).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, "error retrieving users")
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "error reading body")
			return
		}

		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "error parsing user request")
			return
		}

		err = db.Model(&models.User{}).Save(&user).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, "error creating user")
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
