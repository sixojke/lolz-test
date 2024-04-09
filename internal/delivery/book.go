package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sixojke/lolz-test/internal/domain"
)

func (h *Handler) bookCreate(c *gin.Context) {
	var input domain.BookCreateInp
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, errorResponse{
			Code:    400,
			Message: err.Error(),
		})

		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, errorResponse{
			Code:    422,
			Message: err.Error(),
		})

		return
	}

	if err := h.service.Book.Create(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errorResponse{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) bookGetById(c *gin.Context) {
	var input domain.BookGetByIdInp
	input.Id = c.Param("id")

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, errorResponse{
			Code:    422,
			Message: err.Error(),
		})

		return
	}

	book, err := h.service.Book.GetById(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errorResponse{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) booksGetByGenre(c *gin.Context) {
	var input domain.BooksGetByGenreInp
	input.Genre = c.Query("genre")

	input.Limit = processIntParam(c, "limit")
	input.Offset = processIntParam(c, "offset")
	input.Validate(h.config.Books.Limit, h.config.Books.Offset)

	books, err := h.service.Book.GetByGenre(&input)
	if err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, errorResponse{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) bookDelete(c *gin.Context) {
	var input domain.BookDeleteInp
	input.Id = c.Param("id")

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, errorResponse{
			Code:    422,
			Message: err.Error(),
		})

		return
	}

	if err := h.service.Book.Delete(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errorResponse{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) booksSearch(c *gin.Context) {
	var input domain.BooksSearchInp
	input.String = c.Query("string")

	input.Limit = processIntParam(c, "limit")
	input.Offset = processIntParam(c, "offset")

	books, err := h.service.Book.Search(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errorResponse{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, books)
}

func processIntParam(c *gin.Context, paramName string) int {
	num, _ := strconv.Atoi(c.Query(paramName))
	return num
}
