### monitor_server
monitor_server vps面板监控后端api服务

使用说明（源码）:
```shell
1. 拉取代码
2. 进入目录 cd monitor_server
3. 配置好ssl证书，证书放到monitor_server目录下，证书命名 server.cer，密钥为 server.key
4. 执行 go run main.go -p 80004 -secret-key=aaaaaa 启动程序。 -p 设置端口号，-secret-key为app端访问时需要的密钥，自行生成，为了安全起见不要设置过于简单
5. 也可以执行 go build -o monitor_server main.go 编译生成 monitor_server可以执行文件，执行 monitor_server -p 80004 -secret-key=aaaaaa 启动程序
注意，因为app必须要https，上面证书设置部分是必须的。

```