package entity

type DiskIOCounter struct {
	Name       string `json:"name"`
	ReadBytes  uint64 `json:"readBytes"`
	WriteBytes uint64 `json:"writeBytes"`
	ReadSpeed  string `json:"readSpeed"`
	WriteSped  string `json:"writeSpeed"`
}
