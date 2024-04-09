package delivery

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func newErrorResponse(c *gin.Context, statusCode int, r errorResponse) {
	log.Error(r)
	c.AbortWithStatusJSON(statusCode, r)
}
