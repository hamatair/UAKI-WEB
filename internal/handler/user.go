package handler

import (
	"UAKI-WEB/entity"
	"UAKI-WEB/model"
	"UAKI-WEB/pkg/response"
	"errors"
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

func (h *Handler) Login(ctx *gin.Context) {
	param := model.Login{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	token, err := h.Service.UserService.Login(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to login", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success login", token)
}

func (h *Handler) getLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		response.Error(ctx, http.StatusInternalServerError, "failed get login user", errors.New(""))
		return
	}

	response.Success(ctx, http.StatusOK, "get login user", user.(entity.User))
}