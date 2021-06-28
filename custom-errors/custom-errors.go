package custom_errors

import "github.com/gin-gonic/gin"

const (
	ErrorTypeJsonBinding    gin.ErrorType = 100 // JSON Binding
	ErrorTypeJsonValidation gin.ErrorType = 101 // JSON Validation
)
