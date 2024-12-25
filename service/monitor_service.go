package service

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v4/net"
	"monitor_server/entity"
)

type MonitorService struct {
}

/*
*
获取网速
*/
func (m *MonitorService) GetSpeedList() {
	n, _ := net.IOCounters(true)
	fmt.Println(n)
	var speedList []*entity.NetSpeed
	for _, nv := range n {
		fmt.Println(nv.Name)
		fmt.Println(humanize.Bytes(nv.BytesSent))
		fmt.Println(humanize.Bytes(nv.BytesRecv))
		fmt.Println("------------------------------------")

		//currentStats, err := net.NetIOCounters(true)
		netSpeed := &entity.NetSpeed{}
		netSpeed.Name = nv.Name
		speedList = append(speedList, netSpeed)

	}
}
