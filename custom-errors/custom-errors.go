package custom_errors

import "github.com/gin-gonic/gin"

// Bad Request Errors
const (
	ErrorTypeJsonBinding     gin.ErrorType = 100 // JSON Binding
	ErrorTypeJsonValidation  gin.ErrorType = 101 // JSON Validation
	ErrorTypeInvalidArgument gin.ErrorType = 102 // Wrong path parameter, query parameter
	ErrorTypeUnauthorized    gin.ErrorType = 103 // JWT token not present in Header
)

// Database Errors
const (
	ErrorTypeDatabaseError gin.ErrorType = 106
)
