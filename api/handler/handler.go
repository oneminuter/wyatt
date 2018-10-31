package handler

import (
	"net/http"
	"wyatt/api/constant"
	"wyatt/api/view"

	"github.com/gin-gonic/gin"
)

var Page404 = func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, view.SetErr(constant.IllegalRequest))
}
