package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/sixojke/lolz-test/internal/config"
	"github.com/sixojke/lolz-test/internal/service"
)

type Handler struct {
	config  config.HandlerConfig
	service *service.Service
}

func NewHandler(service *service.Service, config config.HandlerConfig) *Handler {
	return &Handler{
		config:  config,
		service: service,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	book := router.Group("/book")
	{
		book.POST("/create", h.bookCreate)
		book.GET("/:id", h.bookGetById)
		book.DELETE("/delete/:id", h.bookDelete)
	}

	books := router.Group("/books")
	{
		books.GET("/list", h.booksGetByGenre)
		books.GET("/search", h.booksSearch)
	}

	genres := router.Group("/genres")
	{
		genres.GET("", h.getGenres)
	}

	return router
}
