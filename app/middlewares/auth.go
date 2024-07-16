package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	e "github.com/mostafababaii/gorest/app/handlers/api/v1/errors"
	"github.com/mostafababaii/gorest/app/handlers/api/v1/response"
	"github.com/mostafababaii/gorest/app/services"
)

func AuthMiddleware(tokenService services.Tokenizer) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := response.NewResponse(c)

		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			r.JsonResponse(http.StatusUnauthorized, e.MISSING_TOKEN, nil)
			c.Abort()
			return
		}

		user, err := tokenService.Validate(clientToken)
		if err != nil {
			r.JsonResponse(http.StatusUnauthorized, e.INVALID_TOKEN, nil)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
