package entity

type ImMessage struct {
	MType   string      `json:"m_type"`
	Data    interface{} `json:"data"`
	NowTime string      `json:"now_time"`
}
