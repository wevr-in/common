package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func GetNatsClient(c *gin.Context) (*nats.Conn, error) {
	sc, e := c.MustGet("sc").(nats.Conn)
	if !e {
		return &nats.Conn{}, errors.New("not found in context")
	}
	return &sc, nil
}
