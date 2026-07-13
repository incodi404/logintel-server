package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"totalPages,omitempty"`
}

func OK(c *gin.Context, message string, data interface{}, meta Meta) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    &meta,
	})
}
