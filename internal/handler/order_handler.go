package handler

import (
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain"
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) MakeOrder(c *gin.Context) {
	json := domain.Order{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}

	userId, _ := c.Get("userId")

	err := h.services.Order.CreateOrder(&json, userId.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewServerGoodResponse("order was added"))
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	json := domain.Order{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}

	err := h.services.Order.UpdateOrderById(&json, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewServerGoodResponse("order was updated"))
}
