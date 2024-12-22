package routers

import (
	"github.com/gin-gonic/gin"
	"monitor_server/api"
	"monitor_server/service"
	"net/http"
)

// 定义中间件函数
func HeaderValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取指定的 Header 参数
		secretKey := c.GetHeader("secret-key")

		// 检查 Header 是否符合要求
		if secretKey == "" || secretKey != service.ConfigSecretKey {
			// 如果验证失败，返回错误响应并终止请求
			//c.JSON(http.StatusBadRequest, gin.H{
			//	"error": "Invalid or missing secret-key",
			//})
			resp := service.ResponseServ.Error(401, "secret-key校验失败")
			c.JSON(http.StatusOK, resp)
			c.Abort() // 终止请求，不会继续传递到后续处理程序
			return
		}

		// 如果验证通过，继续处理下一个中间件或处理程序
		c.Next()
	}
}

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", api.Home)
	apiGroup := r.Group("/api", HeaderValidatorMiddleware())
	apiGroup.GET("/baseinfo", api.BaseInfo)

	return r
}
