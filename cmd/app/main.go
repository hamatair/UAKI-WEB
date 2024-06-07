package main

import (
	"UAKI-WEB/internal/handler"
	"UAKI-WEB/internal/repository"
	"UAKI-WEB/internal/service"
	"UAKI-WEB/pkg/bcrypt"
	"UAKI-WEB/pkg/config"
	"UAKI-WEB/pkg/database"
	"UAKI-WEB/pkg/jwt"
	"UAKI-WEB/pkg/middleware"
	"UAKI-WEB/pkg/websocket"
)

func main() {
	config.LoadEnv()

	jwtAuth := jwt.Init()

	bCrypt := bcrypt.Init()

	db := mysql.ConnectDatabase()

	newRepository := repository.NewRepository(db)

	newService := service.NewService(service.InitParam{
		Repository: newRepository,
		JwtAuth: jwtAuth,
		Bcrypt: bCrypt,
	})

	middleware := middleware.Init(jwtAuth,newService)

	websocket := websocket.Init()

	newHandler := handler.NewHandler(newService, middleware, websocket)

	mysql.Migration(db)

	newHandler.EndPoint()
}