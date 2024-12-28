package service

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
	"monitor_server/entity"
	"sort"
	"strings"
	"time"
)

type MonitorService struct {
}

/*
*
获取网速
*/
func (m *MonitorService) GetSpeedList() (speedList []*entity.NetSpeed) {
	netPre, _ := net.IOCounters(true)
	var speedRecord = make(map[string]net.IOCountersStat)
	for _, nv := range netPre {
		speedRecord[nv.Name] = nv

	}
	time.Sleep(1 * time.Second)
	netCurrent, _ := net.IOCounters(true)
	for _, nc := range netCurrent {
		netSpeed := &entity.NetSpeed{
			Name:          nc.Name,
			UploadSpeed:   humanize.IBytes(nc.BytesSent-speedRecord[nc.Name].BytesSent) + "/s",
			DownloadSpeed: humanize.IBytes(nc.BytesRecv-speedRecord[nc.Name].BytesRecv) + "/s",
		}
		speedList = append(speedList, netSpeed)
	}
	return speedList
}

/*
*
获取磁盘空间使用信息
*/
func (m *MonitorService) GetDiskUsage() (diskUsageList []entity.DiskUsage) {
	partInfos, _ := disk.Partitions(false)
	for _, partInfo := range partInfos {
		if strings.Contains(partInfo.Device, "loop") == false {
			diskUsage := entity.DiskUsage{}
			diskUsage.Device = partInfo.Device
			diskUsage.MountPoint = partInfo.Mountpoint
			diskUsage.Fstype = partInfo.Fstype
			diskUsage.UsageStat, _ = disk.Usage(partInfo.Mountpoint)
			diskUsage.DiskTotalStr = humanize.IBytes(diskUsage.UsageStat.Total)
			diskUsage.DiskFreeStr = humanize.IBytes(diskUsage.UsageStat.Free)
			diskUsage.DiskUsedStr = humanize.IBytes(diskUsage.UsageStat.Used)
			diskUsage.UsedPercent = diskUsage.UsageStat.UsedPercent
			diskUsageList = append(diskUsageList, diskUsage)
		}
	}
	return diskUsageList
}

func (m *MonitorService) GetDiskIOCounters() (diskIOCounters []*entity.DiskIOCounter) {
	var diskIoCounterRecords = make(map[string]disk.IOCountersStat)

	countersPre, _ := disk.IOCounters()
	for k, countersStatPre := range countersPre {
		diskIoCounterRecords[k] = countersStatPre
	}
	time.Sleep(1 * time.Second)
	countersCurrent, _ := disk.IOCounters()
	for key, countersStat := range countersCurrent {
		if strings.Contains(key, "loop") == false {
			diskIOCounter := &entity.DiskIOCounter{}
			diskIOCounter.Name = key
			diskIOCounter.ReadBytes = countersStat.ReadBytes - diskIoCounterRecords[key].ReadBytes
			diskIOCounter.WriteBytes = countersStat.WriteBytes - diskIoCounterRecords[key].WriteBytes
			diskIOCounter.ReadSpeed = humanize.IBytes(diskIOCounter.ReadBytes)
			diskIOCounter.WriteSped = humanize.IBytes(diskIOCounter.WriteBytes)
			diskIOCounters = append(diskIOCounters, diskIOCounter)
		}

	}
	return diskIOCounters
}

/*
*
获取cpu占用前20进程
*/
func (m *MonitorService) GetTopProcessList() (topProcessInfos []entity.ProcessInfo) {
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
	return processInfos[:20]

}
