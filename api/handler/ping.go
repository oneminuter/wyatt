package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var Ping = func(ctx *gin.Context) {
	name, _ := os.Hostname()
	ctx.JSON(http.StatusOK, name)
}
