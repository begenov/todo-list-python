package http

import (
	v1 "github.com/begenov/my-project/backend/internal/delivery/http/v1"
	"github.com/begenov/my-project/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.init(router)

	return router
}

func (h *Handler) init(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.service)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
