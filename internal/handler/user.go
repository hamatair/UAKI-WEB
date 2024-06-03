package handler

import (
	"UAKI-WEB/model"
	"UAKI-WEB/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	param := model.RegisterUser{}

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	newUser, err := h.Service.UserService.RegisterUser(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to register new user", err)
		return
	}

	response.Success(c, http.StatusCreated, "success create new user", newUser)
}