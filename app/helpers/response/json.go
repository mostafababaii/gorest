package response

import (
	"github.com/gin-gonic/gin"
	e "github.com/mostafababaii/gorest/app/helpers/errors"
)

type Response struct {
	C *gin.Context
}

func NewResponse(c *gin.Context) Response {
	return Response{C: c}
}

type ResponseBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) JsonResponse(statusCode, errorCode int, data interface{}) {
	r.C.JSON(statusCode, ResponseBody{
		Code:    errorCode,
		Message: e.GetMessage(errorCode),
		Data:    data,
	})
}
