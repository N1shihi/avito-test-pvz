package handler

import (
	"avito-test-pvz/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserSrv interface {
	CreateUser(user *models.RegisterRequest) (*models.User, error)
}

func (h *Handler) CreateUser(c *gin.Context) {
	log.SetPrefix("Handler Create User")
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	User, err := h.userSrv.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	resp := User.CreateRegisterResponse()
	c.JSON(http.StatusCreated, resp)
}
