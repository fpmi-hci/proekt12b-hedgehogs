package handler

import (
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(h.preflight)
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	router.GET("/api/books/g", h.getAllBooks)
	router.GET("/api/books/author", h.getBookByAuthor)
	router.GET("/api/books/publisher", h.getBookByPublisher)
	api := router.Group("/api", h.userIdentity)
	{
		books := api.Group("/books")
		{
			books.POST("/p", h.createBook)
			books.POST("/p/:id", h.addBookToCart)
			books.GET("/:id", h.getBookById)
			books.PUT("/:id", h.updateBookById)
			books.DELETE("/:id", h.DeleteBookById)
			books.GET("/cart", h.getAllBooksFromCart)
		}

		orders := api.Group("/orders")
		{
			orders.POST("", h.MakeOrder)
			orders.PUT("/:id", h.UpdateOrder)
		}
	}

	return router
}
