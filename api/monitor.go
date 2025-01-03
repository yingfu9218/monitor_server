package api

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"monitor_server/entity"
	"monitor_server/service"
	"net/http"
	"time"
)

func Check(c *gin.Context) {
	resp := service.ResponseServ.Success(true)
	c.JSON(http.StatusOK, resp)
}

func BaseInfo(c *gin.Context) {
	cpuPercent, _ := cpu.Percent(500*time.Millisecond, false)
	cpuPercentStr := fmt.Sprintf("%f", cpuPercent)
	fmt.Printf(cpuPercentStr)
	v, _ := mem.VirtualMemory()
	h, _ := host.Info()
	cpuInfoDetail, _ := cpu.Info()
	cpuino := entity.CpuInfo{
		Cores:     len(cpuInfoDetail),
		ModelName: cpuInfoDetail[0].ModelName,
	}
	baseInfo := entity.BaseInfo{
		HostInfo:        h,
		CpuInfo:         cpuino,
		CpuUsed:         cpuPercent,
		MemInfo:         v,
		MemTotalStr:     humanize.IBytes(v.Total),
		MemAvailableStr: humanize.IBytes(v.Available),
		MemUsedStr:      humanize.IBytes(v.Used),
		MemFreeStr:      humanize.IBytes(v.Free),
	}
	resp := service.ResponseServ.Success(baseInfo)
	c.JSON(http.StatusOK, resp)
}
func DiskUsage(c *gin.Context) {
	monitorService := &service.MonitorService{}
	distusageList := monitorService.GetDiskUsage()
	resp := service.ResponseServ.Success(distusageList)
	c.JSON(http.StatusOK, resp)
}
