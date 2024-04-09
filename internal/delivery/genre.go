package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getGenres(c *gin.Context) {
	genres, err := h.service.Genre.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errorResponse{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, Response{Data: genres})
}
