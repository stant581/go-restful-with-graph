package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stant581/go-restful/models"
	"github.com/stant581/go-restful/services"
)

type UserHandler struct {
	UserService *services.UserService
}

func (handler *UserHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := handler.UserService.GetUserById(userID)
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

func (handler *UserHandler) GetAccountById(ctx *gin.Context) {
	id := ctx.Param("id")
	accountId,err := strconv.Atoi(id)
	if err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error" :"invalid account ID"})
		return
	}
	account, err := handler.UserService.GetAccountById(accountId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func (handler *UserHandler) CreateAccount(ctx *gin.Context) {
	var accountInput struct {
		AccountName string `json:"account_name"`
		Status      bool   `json:"status"`
	}
	if err := ctx.ShouldBindJSON(&accountInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	account := models.Account{
		AccountName: accountInput.AccountName,
		Status:      accountInput.Status,
	}

	createdAccount, err := handler.UserService.CreateAccount(account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createdAccount)	
}

func (handler *UserHandler) GetUserProfileById(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	userProfile, err := handler.UserService.GetUserProfileById(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userProfile)
}

func (handler *UserHandler) CreateUserProfile(ctx *gin.Context) {
	var userProfileInput struct {
		User    *models.User    `json:"user"`
		Account *models.Account `json:"account"`
	}
	if err := ctx.ShouldBindJSON(&userProfileInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	userProfile := models.UserProfile{
		User:    *userProfileInput.User,
		Account: *userProfileInput.Account,
	}

	createdUserProfile, err := handler.UserService.CreateUserProfile(userProfile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createdUserProfile)
}
