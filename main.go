package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/stant581/go-restful/handlers"
	"github.com/stant581/go-restful/repositories"
	"github.com/stant581/go-restful/services"
)

func main() {
	r := gin.Default()

	repo := repositories.NewUserRepository()
	accountRepo := repositories.NewAccountRepository()
	servic := &services.UserService{UserRepo: repo, AccountRepo: accountRepo}
	handler := &handlers.UserHandler{UserService: servic}

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})
	r.GET("/users/:id", handler.GetUserById)
	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id/profile", handler.GetUserProfileById)
	r.POST("/users/profile", handler.CreateUserProfile)
	r.GET("/account/:id",handler.GetAccountById)
	r.POST("/account", handler.CreateAccount)
	r.Run()
}
