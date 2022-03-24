package common

import (
	v1 "goer/app/http/controllers/v1"

	"github.com/gin-gonic/gin"
	"github.com/goer-project/goer/response"
)

type CommonController struct {
	v1.BaseController
}

// Ping
// @BasePath  /
// PingExample godoc
// @Summary   ping server
// @Schemes
// @Description  do ping
// @Tags         Common
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "Ping"
// @Router       /common/ping [get]
func (common *CommonController) Ping(c *gin.Context) {
	response.Msg(c, "pong")
}
