package controllers

import (
	"auth-service/dto"
	"auth-service/models"
	"auth-service/utils"

	"github.com/gin-gonic/gin"
)

func LoginAdmin(c *gin.Context) {
	var loginReq dto.AdminDto

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.Error(gin.Error{})
		return
	}

	if loginReq.Email == "" || loginReq.Password == "" {
		c.Error(utils.ErrorResponse(401, models.ErrorCode[400], "Email and password are required"))
	}
}
