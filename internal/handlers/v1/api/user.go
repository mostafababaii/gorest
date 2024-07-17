package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mostafababaii/gorest/internal/models"
	"github.com/mostafababaii/gorest/internal/services"
	e "github.com/mostafababaii/gorest/pkg/errors"
	"github.com/mostafababaii/gorest/pkg/response"
	"log"
	"net/http"
	"os"
)

type UserHandler struct {
	logger *log.Logger
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (uh *UserHandler) Profile(c *gin.Context) {
	r := response.NewResponse(c)

	user, ok := c.Get("user")
	if !ok {
		r.JsonResponse(http.StatusInternalServerError, e.ERROR, nil)
		uh.logger.Println(errors.New("error on fetching user from request"))
		return
	}

	user, err := services.UserService.FindByID(user.(*models.User).ID)
	if err != nil {
		r.JsonResponse(http.StatusNotFound, e.USER_NOT_FOUND, nil)
		uh.logger.Println(err)
		return
	}

	r.JsonResponse(http.StatusOK, e.SUCCESS, user)
}
