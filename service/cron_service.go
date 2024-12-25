package service

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v4/cpu"
	"monitor_server/entity"
	"monitor_server/lib"
	"time"
)

type CronService struct {
}

var CronServ *CronService = &CronService{}

func (c *CronService) Start() {
	fmt.Println("cron start")
	go c.exec()
}
func (c *CronService) exec() {
	fmt.Println("cron exec")
	ct := cron.New()
	ct.AddFunc("@every 5s", func() {
		fmt.Println("cron Task executed")
		fmt.Println("当前会话数：")
		fmt.Println(len(WebSocketManagerServ.GetConnections()))

		for _, c := range WebSocketManagerServ.GetConnections() {
			fmt.Println("connid:" + c.ID)

			imMessage := &entity.ImMessage{
				MType: "text",
				Data:  "响应：11111",
			}
			mesage, _ := json.Marshal(imMessage)
			c.Conn.WriteMessage(1, mesage)
			cpuPercnt, _ := cpu.Percent(5*time.Second, false)
			imMessage = &entity.ImMessage{
				MType:   "cpu_percent",
				Data:    cpuPercnt[0],
				NowTime: lib.GetTimeNow(),
			}
			mesage, _ = json.Marshal(imMessage)
			c.Conn.WriteMessage(1, mesage)
		}
	})
	ct.Start()
	select {}
}
