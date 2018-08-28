package handler

import (
	"net/http"

	"strings"
	"wyatt/util"

	"github.com/gin-gonic/gin"
)

var UserInfo = func(ctx *gin.Context) {
	var token = ctx.Query("token")
	if strings.TrimSpace(token) == "" {
		token = util.GetToken()
	}

	ctx.JSON(http.StatusOK, struct {
		Token string
	}{
		Token: token,
	})
}
