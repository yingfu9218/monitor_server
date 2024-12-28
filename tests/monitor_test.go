package tests

import (
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
	"monitor_server/entity"
	"monitor_server/service"
	"sort"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	cpuPercent, _ := cpu.Percent(500*time.Millisecond, false)
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

func speedTime(dfn func()) {
	start := time.Now()
	dfn()
	duration := time.Since(start)
	// 输出执行时长
	fmt.Printf(" \nFunction execution time: %v\n", duration)
}

func TestSpeedTest(t *testing.T) {
	speedTime(func() {
		cpuPercent, _ := cpu.Percent(500*time.Millisecond, false)
		cpuPercentStr := fmt.Sprintf("%f", cpuPercent)
		fmt.Printf(cpuPercentStr)
	})

	speedTime(func() {
		cpuInfo, _ := cpu.Info()
		cpuInfoStr, _ := json.Marshal(cpuInfo)
		fmt.Printf(string(cpuInfoStr))
	})
	speedTime(func() {
		v, _ := mem.VirtualMemory()
		vStr, _ := json.Marshal(v)
		fmt.Printf(string(vStr))
		h, _ := host.Info()
		hStr, _ := json.Marshal(h)
		fmt.Printf(string(hStr))
	})

}
func TestDisk(t *testing.T) {
	d, _ := disk.IOCounters()
	fmt.Println(d)
	d2, _ := disk.Partitions(false)
	fmt.Println(d2)
	d3, _ := disk.Usage("/")
	fmt.Println(d3)
}

func TestProcess(t *testing.T) {
	// 获取所有进程
	procs, err := process.Processes()
	if err != nil {
		fmt.Println("Error getting processes:", err)
		return
	}

	var processInfos []entity.ProcessInfo

	// 遍历所有进程并获取CPU占用率
	for _, p := range procs {
		name, err := p.Name()
		if err != nil {
			continue
		}
		cpuPercent, err := p.CPUPercent()
		memPercent, err := p.MemoryPercent()
		a, _ := p.Cmdline()
		fmt.Println(a)
		if err != nil {
			continue
		}

		processInfos = append(processInfos, entity.ProcessInfo{
			Pid:        p.Pid,
			Name:       name,
			CpuPercent: cpuPercent,
			MemPercent: memPercent,
		})
	}

	// 按照CPU占用率降序排序
	sort.Slice(processInfos, func(i, j int) bool {
		return processInfos[i].CpuPercent > processInfos[j].CpuPercent
	})

	// 输出前十个进程
	fmt.Println("Top 10 processes by CPU usage:")
	for i := 0; i < 10 && i < len(processInfos); i++ {
		fmt.Printf("%d: %s (PID: %d) - CPU Usage: %.2f%%\n", i+1, processInfos[i].Name, processInfos[i].Pid, processInfos[i].CpuPercent)
	}

	// 延时，等待CPU占用率变化（可选）
	time.Sleep(2 * time.Second)

}

func TestNet(t *testing.T) {
	n, _ := net.IOCounters(true)
	fmt.Println(n)

	for _, nv := range n {
		fmt.Println(nv.Name)
		fmt.Println(humanize.Bytes(nv.BytesSent))
		fmt.Println(humanize.Bytes(nv.BytesRecv))
		fmt.Println("------------------------------------")
	}

}

func TestGetSpeedTestList(t *testing.T) {
	monitorService := &service.MonitorService{}
	speedList := monitorService.GetSpeedList()
	fmt.Println(speedList)
}
func TestGetDiskUsage(t *testing.T) {
	monitorService := &service.MonitorService{}
	distusageList := monitorService.GetDiskUsage()
	fmt.Println(distusageList)
}
func TestGetDiskIOCounters(t *testing.T) {
	monitorService := &service.MonitorService{}
	diskIoCounters := monitorService.GetDiskIOCounters()
	fmt.Println(diskIoCounters)
}
