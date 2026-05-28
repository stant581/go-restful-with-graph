package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/stant581/go-restful/models"
	"github.com/stant581/go-restful/services"
)

type UserHandler struct {
	UserService *services.UserService
	schema      *graphql.Schema
	initialized bool
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
	accountId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
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

// GraphQL Handlers

func (handler *UserHandler) initSchema() *graphql.Schema {
	// Define User type
	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"userId": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"userName": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"accountId": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"account": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "Account",
					Fields: graphql.Fields{
						"accountId": &graphql.Field{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"accountName": &graphql.Field{
							Type: graphql.NewNonNull(graphql.String),
						},
						"status": &graphql.Field{
							Type: graphql.NewNonNull(graphql.Boolean),
						},
					},
				}),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*models.User); ok {
						account, err := handler.UserService.GetAccountById(user.AccountId)
						if err != nil {
							return nil, nil
						}
						return account, nil
					}
					return nil, nil
				},
			},
		},
	})

	// Define Account type
	accountType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Account",
		Fields: graphql.Fields{
			"accountId": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"accountName": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"status": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
		},
	})

	// Define UserProfile type
	userProfileType := graphql.NewObject(graphql.ObjectConfig{
		Name: "UserProfile",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: graphql.NewNonNull(userType),
			},
			"account": &graphql.Field{
				Type: accountType,
			},
		},
	})

	// Define Query type
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok {
						return nil, nil
					}
					user, err := handler.UserService.GetUserById(id)
					if err != nil {
						return nil, nil
					}
					return &user, nil
				},
			},
			"account": &graphql.Field{
				Type: accountType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok {
						return nil, nil
					}
					account, err := handler.UserService.GetAccountById(id)
					if err != nil {
						return nil, nil
					}
					return &account, nil
				},
			},
			"userProfile": &graphql.Field{
				Type: userProfileType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok {
						return nil, nil
					}
					profile, err := handler.UserService.GetUserProfileById(id)
					if err != nil {
						return nil, nil
					}
					return &profile, nil
				},
			},
		},
	})

	// Define Mutation type
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"userName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"accountId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					userName, _ := p.Args["userName"].(string)
					accountId, _ := p.Args["accountId"].(int)
					user := models.User{
						UserName:  userName,
						AccountId: accountId,
					}
					createdUser, err := handler.UserService.CreateUser(user)
					if err != nil {
						return nil, err
					}
					return &createdUser, nil
				},
			},
			"createAccount": &graphql.Field{
				Type: accountType,
				Args: graphql.FieldConfigArgument{
					"accountName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"status": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					accountName, _ := p.Args["accountName"].(string)
					status, _ := p.Args["status"].(bool)
					account := models.Account{
						AccountName: accountName,
						Status:      status,
					}
					createdAccount, err := handler.UserService.CreateAccount(account)
					if err != nil {
						return nil, err
					}
					return &createdAccount, nil
				},
			},
		},
	})

	// Create schema
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return &schema
}

func (handler *UserHandler) GraphQLHandler(ctx *gin.Context) {
	if !handler.initialized {
		handler.schema = handler.initSchema()
		handler.initialized = true
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var req struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.Unmarshal(body, &req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	params := graphql.Params{
		Schema:         *handler.schema,
		RequestString:  req.Query,
		OperationName:  req.OperationName,
		VariableValues: req.Variables,
	}

	result := graphql.Do(params)

	if len(result.Errors) > 0 {
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (handler *UserHandler) PlaygroundHandler(ctx *gin.Context) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>GraphQL Playground</title>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<script src="https://unpkg.com/graphql-playground-umd/build/static/js/middleware.js"></script>
	</head>
	<body>
		<div id="root"></div>
		<script>
			window.addEventListener('load', function (event) {
				GraphQLPlayground.init(document.getElementById('root'), {
					endpoint: '/graphql',
					subscriptionEndpoint: '/graphql',
				})
			})
		</script>
	</body>
	</html>
	`
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, html)
}
