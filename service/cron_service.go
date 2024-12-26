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
			//网速
			monitorService := &MonitorService{}
			netSpeedLlist := monitorService.GetSpeedList()
			imMessage = &entity.ImMessage{
				MType:   "net_speed_list",
				Data:    netSpeedLlist,
				NowTime: lib.GetTimeNow(),
			}
			netSpeedMesage, _ := json.Marshal(imMessage)
			//磁盘读写速率
			diskIOCounters := monitorService.GetDiskIOCounters()
			imMessage = &entity.ImMessage{
				MType:   "disk_io_counters",
				Data:    diskIOCounters,
				NowTime: lib.GetTimeNow(),
			}
			diskUsageMesage, _ := json.Marshal(imMessage)

			//推送消息
			for _, c := range WebSocketManagerServ.GetConnections() {
				fmt.Println("connid:" + c.ID)
				c.Conn.WriteMessage(1, HelloMesage)
				c.Conn.WriteMessage(1, cpuPercentMesage)
				c.Conn.WriteMessage(1, menDetailMesage)
				c.Conn.WriteMessage(1, netSpeedMesage)
				c.Conn.WriteMessage(1, diskUsageMesage)

			}
		}

	})
	ct.Start()
	select {}
}
