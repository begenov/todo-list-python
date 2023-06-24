package v1

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message interface{} `json:"message"`
}

func newResponse(ctx *gin.Context, code int, message ...interface{}) {
	log.Println(message...)
	ctx.JSON(code, Response{Message: message})
}
