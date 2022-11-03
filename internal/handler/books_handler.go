package handler

import (
	"net/http"
	"strconv"

	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain"
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain/responses"
	"github.com/gin-gonic/gin"
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

func (h *Handler) addBookToCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	userId, _ := c.Get("userId")

	err = h.services.AddBookToCart(id, userId.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewServerGoodResponse("book was added"))
}

func (h *Handler) getAllBooks(c *gin.Context) {

	filters := domain.BookFilter{
		Name:      c.Query("title"),
		Category:  c.Query("category"),
		Author:    c.Query("author"),
		Publisher: c.Query("publisher"),
		Sort:      c.Query("sort"),
	}

	books, err := h.services.GetAllBooks(&filters)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) getAllBooksFromCart(c *gin.Context) {

	userId, _ := c.Get("userId")

	books, err := h.services.GetBookFromCartByUserId(userId.(int))
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
	userId, _ := c.Get("userId")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	err = h.services.DeleteBookById(id, userId.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.NewServerGoodResponse("Books was deleted"))
}
