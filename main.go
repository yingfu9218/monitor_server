package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"monitor_server/routers"
	"monitor_server/service"
	"net/http"
	"os"
	"strconv"
	"time"
)

var port int
var secretKey string

func main() {

	flag.IntVar(&port, "p", 8004, "listen port ,default :8004")
	flag.StringVar(&secretKey, "secret-key", "", "secret-key ,default :")
	flag.Parse()
	if secretKey == "" {
		fmt.Printf("secret-key 不能为空，请输入相关的密钥")
		os.Exit(0)
	}
	service.ConfigSecretKey = secretKey
	gin.SetMode("debug")
	//启动定时器
	service.CronServ.Start()
	routersInit := routers.InitRouter()
	readTimeout := 30 * time.Second
	writeTimeout := 30 * time.Second
	endPoint := ":" + strconv.Itoa(port)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	log.Printf("服务器已启动成功，请在浏览器访问 http://localhost:%d   调试\n", port)

	server.ListenAndServe()

}
