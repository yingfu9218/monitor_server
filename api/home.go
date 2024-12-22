package api

import (
	"github.com/gin-gonic/gin"
	"monitor_server/service"
	"net/http"
)

func Home(c *gin.Context) {
	resp := service.ResponseServ.Success(true)
	c.JSON(http.StatusOK, resp)
}
