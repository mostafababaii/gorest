package middlewares

import (
	"github.com/mostafababaii/gorest/internal/interfaces"
	e "github.com/mostafababaii/gorest/pkg/errors"
	"github.com/mostafababaii/gorest/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenService interfaces.Tokenizer) gin.HandlerFunc {
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
