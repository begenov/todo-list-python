package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initRoutersAccounts(api *gin.RouterGroup) {
	accounts := api.Group("/accounts")
	{
		accounts.POST("/sign-up", h.accountsSignUp)
		accounts.POST("/sign-in", h.accountsSignIn)
		accounts.POST("/refresh-token", h.accountsRefreshToken)
		authentocated := accounts.Group("/")
		{
			authentocated.GET("/me", h.accountsMe)
			authentocated.POST("/image", h.accountsImage)
			authentocated.DELETE("/image", h.accountsImageDelete)
			authentocated.PUT("/details", h.accountsDetails)
		}

	}
}

type signUp struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (h *Handler) accountsSignUp(c *gin.Context) {
	var inp signUp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, Response{"StatusOK"})
}

type signIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) accountsSignIn(c *gin.Context) {
	var inp signIn
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, Response{"StatusOK"})
}

func (h *Handler) accountsRefreshToken(c *gin.Context) {
}

func (h *Handler) accountsMe(c *gin.Context) {
}

func (h *Handler) accountsImage(c *gin.Context) {
}

func (h *Handler) accountsImageDelete(c *gin.Context) {
}

func (h *Handler) accountsDetails(c *gin.Context) {
}
