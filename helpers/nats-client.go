package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
)

func GetNatsClient(c *gin.Context) (*stan.Conn, error) {
	sc, e := c.MustGet("sc").(stan.Conn)
	if !e {
		return nil, errors.New("not found in context")
	}
	return &sc, nil
}
