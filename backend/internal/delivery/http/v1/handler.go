package v1

import (
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

func (h *Handler) Init(api *gin.RouterGroup) {
	h.initRoutersAccounts(api)
}
