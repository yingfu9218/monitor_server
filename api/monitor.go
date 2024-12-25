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

func BaseInfo(c *gin.Context) {
	cpuPercent, _ := cpu.Percent(500*time.Millisecond, false)
	cpuPercentStr := fmt.Sprintf("%f", cpuPercent)
	fmt.Printf(cpuPercentStr)
	v, _ := mem.VirtualMemory()
	h, _ := host.Info()
	baseInfo := entity.BaseInfo{
		HostInfo:        h,
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
