package entity

type CpuInfo struct {
	Cores     int    `json:"cores"`
	ModelName string `json:"modelName"`
}
