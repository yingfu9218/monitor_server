package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"monitor_server/api"
	"monitor_server/service"
	"net/http"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", api.Home)
	r.GET("/ws", WebSocketHandler)
	apiGroup := r.Group("/api", HeaderValidatorMiddleware())
	apiGroup.GET("/baseinfo", api.BaseInfo)

	return r
}

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

/*
*
websocket
*/
func WebSocketHandler(c *gin.Context) {
	token := c.DefaultQuery("token", "")
	fmt.Println(token)
	// 获取WebSocket连接
	ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功")
	connID := service.WebSocketManagerServ.CreateSessionId()
	service.WebSocketManagerServ.AddConnection(connID, ws)
	defer func() {
		// 关闭WebSocket连接
		ws.Close()
		fmt.Println("关闭连接")
	}()

	// 处理WebSocket消息
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println("messageType:", messageType)
		fmt.Println("p:", string(p))
		// 输出WebSocket消息内容
		p = append([]byte("server response:"), p...)
		err = ws.WriteMessage(messageType, p)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}
	ws.Close()
	service.WebSocketManagerServ.RemoveConnection(connID)
	fmt.Println(service.WebSocketManagerServ.GetConnections())

}
