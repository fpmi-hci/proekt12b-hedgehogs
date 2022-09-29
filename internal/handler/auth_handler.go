package handler

import (
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain"
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain/responses"
	"github.com/gin-gonic/gin"

	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	json := domain.User{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	simplePassword := json.PasswordHash
	user, err := h.services.CreateUser(&json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}

	token, err := h.services.GenerateToken(json.Login, simplePassword)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

type signInInput struct {
	Login        string `db:"login" json:"login" binding:"required"`
	PasswordHash string `db:"password_hash" json:"passwordHash" binding:"required,gte=8"`
}

func (h *Handler) signIn(c *gin.Context) {
	json := signInInput{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewServerUnauthorizedResponse(err.Error()))
		return
	}
	token, err := h.services.GenerateToken(json.Login, json.PasswordHash)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}

	user, err := h.services.GetUserByUsername(json.Login)
	user.PasswordHash = ""

	c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})

}
