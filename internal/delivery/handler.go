package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/sixojke/lolz-test/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	book := router.Group("/book")
	{
		book.POST("/create", h.bookCreate)
		book.GET("/:id", h.bookGetById)
		book.GET("/search")
		book.DELETE("/delete", h.bookDelete)
	}

	books := router.Group("/books")
	{
		books.GET("/list")
	}

	return router
}
