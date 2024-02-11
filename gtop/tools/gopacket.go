package tools

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// win：wireshark - 底层 npcap
// linux：tcpdump - 底层 libpcap
// go：github.com/google/gopacket

var (
	VarPacketDeviceName string
)

func packetServiceCommand(_ *cobra.Command, args []string) error {

	//获取所有的网络设备
	devs, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	for _, dev := range devs {
		fmt.Println("网卡：", dev.Name, dev.Addresses)
	}

	/*
		pcap.OpenLive() 函数打开一个网络设备并创建一个handler对象，该对象用于抓包操作
		1.网卡名称
		2.传输大小
		3.混合模式
		4.抓取间隔时间
	*/
	handler, err := pcap.OpenLive(VarPacketDeviceName, 1024, false, time.Second*3)
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()

	//创建source 对象，响应抓包通道，接收抓包数据
	source := gopacket.NewPacketSource(handler, handler.LinkType())

	//遍历通道
	for packet := range source.Packets() {

		//过滤必须 osi七层中的 transport 网络级别才可以输出
		if tcpPlayer := packet.TransportLayer(); tcpPlayer != nil {

			tcp4, ok := tcpPlayer.(*layers.TCP)
			if !ok {
				continue
			}
			fmt.Println(tcp4.DstPort, tcp4.SrcPort)
			//DstPort 请求，SrcPort 响应
			if tcp4.DstPort == 8111 || tcp4.SrcPort == 8111 {
				fmt.Println(string(tcp4.Payload))
			}
		}
	}

	return nil
}
