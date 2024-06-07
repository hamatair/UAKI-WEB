package handler

import (
	"UAKI-WEB/internal/service"
	"UAKI-WEB/pkg/middleware"
	"UAKI-WEB/pkg/websocket"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service    *service.Service
	Router     *gin.Engine
	Middleware middleware.Interface
	Websocket  websocket.Interface
}

func NewHandler(service *service.Service, middleware middleware.Interface, websocket websocket.Interface) *Handler {
	return &Handler{
		Service:    service,
		Router:     gin.Default(),
		Middleware: middleware,
		Websocket:  websocket,
	}
}

func (h *Handler) EndPoint() {
	h.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	go h.Websocket.RunWebsocket()

	v1 := h.Router.Group("/v1")

	v1.POST("user/register", h.RegisterUser)
	v1.GET("user/get-user", h.Middleware.AuthenticateUser, h.getLoginUser)
	v1.POST("user/login", h.Login)

	v1.GET("websocket", h.Websocket.ServeWS)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	h.Router.Run(fmt.Sprintf(":%s", port))
}
