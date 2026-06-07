package server

import (
	"net/http"

	"github.com/clevanilson/cs-url-shortner/pkg/c_errors"
	"github.com/gin-gonic/gin"
)

func HandleError(err error, c *gin.Context) {
	domainErr, ok := err.(*c_errors.DomainError)
	if ok {
		c.JSON(http.StatusBadRequest, formatError(domainErr))
		return
	}
	notFoundErr, ok := err.(*c_errors.NotFoundError)
	if ok {
		c.JSON(http.StatusNotFound, formatError(notFoundErr))
		return
	}
	if ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Try again later"})
		return
	}
}

func formatError(err c_errors.CError) *gin.H {
	return &gin.H{"message": err.Error()}
}
