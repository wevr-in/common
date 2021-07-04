package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
)

func NatsApiMiddleware(sc stan.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("sc", sc)
	}
}
