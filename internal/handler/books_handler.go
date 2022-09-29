package handler

import (
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain"
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createBook(c *gin.Context) {
	json := domain.Book{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	userId, _ := c.Get("userId")

	err := h.services.CreateBook(&json, userId.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewServerGoodResponse("book was added"))
}

func (h *Handler) getAllBooks(c *gin.Context) {

	books, err := h.services.GetAllBooks()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) getBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	book, err := h.services.GetBookById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) getBookByAuthor(c *gin.Context) {
	author := c.Query("author")
	book, err := h.services.GetBookByAuthor(author)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) getBookByPublisher(c *gin.Context) {
	publisher := c.Query("publisher")
	book, err := h.services.GetBookByPublisher(publisher)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) updateBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	json := domain.Book{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	err = h.services.UpdateBookById(&json, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.NewServerGoodResponse("Books was updated"))
}

func (h *Handler) DeleteBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	err = h.services.DeleteBookById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.NewServerGoodResponse("Books was deleted"))
}
