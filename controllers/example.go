package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary ping example
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/ping [get]
func (c *Controller) PingExample(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
	return
}
