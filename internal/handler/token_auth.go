package handler

import (
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const authorizationHeader = "Authorization"

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewServerUnauthorizedResponse("empty auth header"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewServerUnauthorizedResponse("invalid auth header"))
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewServerUnauthorizedResponse(err.Error()))
		return
	}

	c.Set("userId", userId)
}
