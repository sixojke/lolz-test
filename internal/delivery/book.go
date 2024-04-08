package delivery

import (
	"net/http"

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
	input.Id = c.Query("id")

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

func (h *Handler) bookDelete(c *gin.Context) {
	var input domain.BookDeleteInp
	input.Id = c.Query("id")

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
