package service

import (
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
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
	ct.AddFunc("@every 1s", func() {
		fmt.Println("cron Task executed")
		fmt.Println("当前会话数：")
		fmt.Println(len(WebSocketManagerServ.GetConnections()))
		//有连接时执行
		if len(WebSocketManagerServ.GetConnections()) > 0 {

			imMessage := &entity.ImMessage{
				MType: "text",
				Data:  "响应：11111",
			}
			HelloMesage, _ := json.Marshal(imMessage)
			//获取cpu使用率
			cpuPercnt, _ := cpu.Percent(2*time.Second, false)
			imMessage = &entity.ImMessage{
				MType:   "cpu_percent",
				Data:    cpuPercnt[0],
				NowTime: lib.GetTimeNow(),
			}
			cpuPercentMesage, _ := json.Marshal(imMessage)
			//内存使用率
			v, _ := mem.VirtualMemory()
			menDetail := entity.MemDetail{
				MemInfo:         v,
				MemTotalStr:     humanize.IBytes(v.Total),
				MemAvailableStr: humanize.IBytes(v.Available),
				MemUsedStr:      humanize.IBytes(v.Used),
				MemFreeStr:      humanize.IBytes(v.Free),
			}
			imMessage = &entity.ImMessage{
				MType:   "mem_datail",
				Data:    menDetail,
				NowTime: lib.GetTimeNow(),
			}
			menDetailMesage, _ := json.Marshal(imMessage)
			//网络
			for _, c := range WebSocketManagerServ.GetConnections() {
				fmt.Println("connid:" + c.ID)
				c.Conn.WriteMessage(1, HelloMesage)
				c.Conn.WriteMessage(1, cpuPercentMesage)
				c.Conn.WriteMessage(1, menDetailMesage)

			}
		}

	})
	ct.Start()
	select {}
}
