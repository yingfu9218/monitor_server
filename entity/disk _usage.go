package entity

import "github.com/shirou/gopsutil/v4/disk"

type DiskUsage struct {
	Device       string          `json:"device"`
	MountPoint   string          `json:"mountPoint"`
	Fstype       string          `json:"fstype"`
	UsageStat    *disk.UsageStat `json:"usageStat"`
	DiskTotalStr string          `json:"diskTotalStr"`
	DiskFreeStr  string          `json:"diskFreeStr"`
	DiskUsedStr  string          `json:"diskUsedStr"`
	UsedPercent  float64         `json:"usedPercent"`
}
