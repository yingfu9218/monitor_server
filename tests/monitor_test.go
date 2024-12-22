package tests

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"testing"
	"time"
)

func TestCpuInfo(t *testing.T) {
	cpuPercent, _ := cpu.Percent(3*time.Second, false)
	cpuPercentStr := fmt.Sprintf("%f", cpuPercent)
	fmt.Printf(cpuPercentStr)
	cpuInfo, _ := cpu.Info()
	cpuInfoStr, _ := json.Marshal(cpuInfo)
	fmt.Printf(string(cpuInfoStr))
	v, _ := mem.VirtualMemory()
	s, _ := json.Marshal(v)
	fmt.Printf(string(s))

	d, _ := disk.Usage("/")
	dstr, _ := json.Marshal(d)
	fmt.Printf(string(dstr))
	h, _ := host.Info()
	hstr, _ := json.Marshal(h)
	fmt.Printf(string(hstr))

	platform, family, version, _ := host.PlatformInformation()
	fmt.Println("platform:", platform)
	fmt.Println("family:", family)
	fmt.Println("version:", version)

	//baseInfo := entity.BaseInfo{
	//	CpuInfo: cpuInfo,
	//	CpuUsed: cpuPercent,
	//	MemUsed: float64(v.Used) / float64(v.Total),
	//}
	ioStats, err := net.IOCounters(false) // 参数 true 表示按每个网络接口分别统计
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, stat := range ioStats {
		fmt.Printf("Name: %s\n", stat.Name)
		fmt.Printf("BytesSent: %d\n", stat.BytesSent)
		fmt.Printf("BytesRecv: %d\n", stat.BytesRecv)
		fmt.Printf("PacketsSent: %d\n", stat.PacketsSent)
		fmt.Printf("PacketsRecv: %d\n", stat.PacketsRecv)
		fmt.Println()
	}

}
