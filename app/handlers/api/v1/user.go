package v1

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	e "github.com/mostafababaii/gorest/app/handlers/api/v1/errors"
	"github.com/mostafababaii/gorest/app/handlers/api/v1/response"
	"github.com/mostafababaii/gorest/app/models"
)

type UserHandler struct {
	User   models.User
	logger *log.Logger
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		User:   models.NewUser(),
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (u *UserHandler) Profile(c *gin.Context) {
	r := response.NewResponse(c)

	requestUser, ok := c.Get("user")
	if !ok {
		r.JsonResponse(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	user, err := u.User.FindByID(requestUser.(*models.User).ID)
	if err != nil {
		r.JsonResponse(http.StatusNotFound, e.USER_NOT_FOUND, nil)
		return
	}

	r.JsonResponse(http.StatusOK, e.SUCCESS, user)
}
