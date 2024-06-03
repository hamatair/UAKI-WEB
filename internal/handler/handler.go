package handler

import (
	"UAKI-WEB/internal/service"
	"UAKI-WEB/pkg/middleware"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
	Router  *gin.Engine
	Middleware middleware.Interface
}

func NewHandler(service *service.Service, middleware middleware.Interface) *Handler {
	return &Handler{
		Service: service,
		Router:  gin.Default(),
		Middleware: middleware,
	}
}

func (h *Handler) EndPoint(){
	h.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	
	v1 := h.Router.Group("/v1")

	v1.POST("user/register", h.RegisterUser)
	v1.GET("user/get-user", )
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	h.Router.Run(fmt.Sprintf(":%s", port))
}