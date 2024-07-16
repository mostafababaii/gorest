package api

import (
	"github.com/mostafababaii/gorest/internal/interfaces"
	"github.com/mostafababaii/gorest/internal/models"
	"github.com/mostafababaii/gorest/internal/services"
	e "github.com/mostafababaii/gorest/pkg/errors"
	"github.com/mostafababaii/gorest/pkg/response"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	User         models.User
	TokenService interfaces.Tokenizer
	logger       *log.Logger
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		User:         models.NewUser(),
		TokenService: services.NewJwtToken(),
		logger:       log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (a *AuthHandler) Register(c *gin.Context) {
	r := response.NewResponse(c)

	var body models.RegisterBody
	if err := c.ShouldBindJSON(&body); err != nil {
		r.JsonResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		a.logger.Println(err)
		return
	}

	var user models.User
	body.MapTo(&user)

	registeredUser, err := a.User.Create(user)
	if err != nil {
		r.JsonResponse(http.StatusInternalServerError, e.ERROR, nil)
		a.logger.Println(err)
		return
	}

	token, err := a.TokenService.Generate(registeredUser)
	if err != nil {
		r.JsonResponse(http.StatusInternalServerError, e.ERROR, nil)
		a.logger.Println(err)
		return
	}

	r.JsonResponse(http.StatusCreated, e.SUCCESS, map[string]any{"token": token})
}

func (a *AuthHandler) Login(c *gin.Context) {
	r := response.NewResponse(c)

	var body models.LoginBody
	if err := c.ShouldBindJSON(&body); err != nil {
		r.JsonResponse(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	var user models.User
	body.MapTo(&user)

	authUser, ok := a.User.Authenticate(user)
	if !ok {
		r.JsonResponse(http.StatusUnauthorized, e.AUTHENTICATION_FAILED, nil)
		return
	}

	token, err := a.TokenService.Generate(authUser)
	if err != nil {
		r.JsonResponse(http.StatusInternalServerError, e.ERROR, nil)
		a.logger.Println(err)
		return
	}

	r.JsonResponse(http.StatusOK, e.SUCCESS, map[string]any{"token": token})
}
