package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stant581/go-restful/handlers"
	"github.com/stant581/go-restful/repositories"
	"github.com/stant581/go-restful/services"
)

func main() {
	r := gin.Default()

	repo := repositories.NewUserRepository()
	servic := &services.UserService{UserRepo: repo}
	handler := &handlers.UserHandler{UserService: servic}

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})
	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		userID, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}
		handler.GetUserById(userID, ctx)
	})
	r.POST("/users", handler.CreateUser)
	r.Run()
}
