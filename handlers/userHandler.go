package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stant581/go-restful/models"
	"github.com/stant581/go-restful/services"
)

type UserHandler struct {
	UserService *services.UserService
}

func (handler *UserHandler) GetUserById(id int, ctx *gin.Context) {
	user, err := handler.UserService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (handler *UserHandler) CreateUser(ctx *gin.Context) {
	var userInput struct {
		UserName  string `json:"user_name"`
		AccountId int    `json:"account_id"`
	}
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	user := models.User{
		UserName:  userInput.UserName,
		AccountId: userInput.AccountId,
	}

	createdUser, err := handler.UserService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createdUser)
}
