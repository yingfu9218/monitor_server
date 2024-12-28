package entity

type ProcessInfo struct {
	Pid        int32   `json:"pid"`
	Name       string  `json:"name"`
	CpuPercent float64 `json:"cpuPercent"`
	MemPercent float32 `json:"memPercent"`
}
