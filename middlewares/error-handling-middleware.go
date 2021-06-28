package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
)

const errResp string = "error"

func ErrorHandlingMiddleware(trans ut.Translator, val *validator.Validate) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("trans", trans)
		c.Set("validator", val)
		c.Next()

		// Checking the custom-errors and setting appropriate error messages in format { "custom-errors": [...] }
		for _, e := range c.Errors {
			switch e.Type {
			case 100:
				var se []string
				se = append(se, e.Err.Error())
				c.JSON(http.StatusBadRequest, gin.H{
					errResp: se,
				})
				break
			case 101:
				var ve error
				errors.As(e.Err, &ve)

				ves := ve.(validator.ValidationErrors)

				var se []string
				for _, ve := range ves {
					println(ve.Translate(trans))
					se = append(se, ve.Translate(trans))
				}
				c.JSON(http.StatusBadRequest, gin.H{
					errResp: se,
				})
				break
			default:
				var se []string
				se = append(se, e.Err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					errResp: se,
				})
			}
		}
	}
}

// Tag added
