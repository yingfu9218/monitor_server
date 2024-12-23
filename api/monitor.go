package api

import (
	"encoding/json"
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
	cpuInfo, _ := cpu.Info()
	cpuInfoStr, _ := json.Marshal(cpuInfo)
	fmt.Printf(string(cpuInfoStr))
	v, _ := mem.VirtualMemory()
	h, _ := host.Info()
	baseInfo := entity.BaseInfo{
		HostInfo:        h,
		CpuInfo:         cpuInfo,
		CpuUsed:         cpuPercent,
		MemInfo:         v,
		MemTotalStr:     humanize.Bytes(v.Total),
		MemAvailableStr: humanize.Bytes(v.Available),
		MemUsedStr:      humanize.Bytes(v.Used),
		MemFreeStr:      humanize.Bytes(v.Free),
	}
	resp := service.ResponseServ.Success(baseInfo)
	c.JSON(http.StatusOK, resp)
}
