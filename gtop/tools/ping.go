package tools

import (
	"errors"
	"fmt"
	probing "github.com/prometheus-community/pro-bing"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"time"
)

/*
	github.com/prometheus-community/pro-bing
	[-c count] [-i interval] [-t timeout] [--privileged] host

	Count：指定要发送的Ping请求的次数。默认值为$-1$，表示无限次数。
	Size：指定每个Ping请求的数据包大小（字节）。默认值为24字节。
	Interval：指定发送Ping请求之间的时间间隔。默认值为1秒。
	Timeout：指定等待Ping响应的超时时间。默认值为100秒。
	Privileged：指定是否以特权模式发送Ping请求。默认值为false，表示以普通用户权限发送。

    # ping google continuously
    ping www.google.com

    # ping google 5 times
    ping -c 5 www.google.com

    # ping google 5 times at 500ms intervals
    ping -c 5 -i 500ms www.google.com

    # ping google for 10 seconds
    ping -t 10s www.google.com

    # Send a privileged raw ICMP ping
    sudo ping --privileged www.google.com

    # Send ICMP messages with a 100-byte payload
    ping -s 100 1.1.1.1
*/

var (
	VarPingCountInt   int
	VarPingInterval   int
	VarPingTimeout    int
	VarPingPrivileged bool
	mustHostErr       = errors.New("the specified host or ip address must be passed")
)

func pingServiceCommand(_ *cobra.Command, args []string) error {

	if len(args) == 0 {
		return mustHostErr
	}
	host := args[0]

	pinger, err := probing.NewPinger(host)
	if err != nil {
		log.Println("ping err:", err)
		return err
	}

	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *probing.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
	}
	pinger.OnDuplicateRecv = func(pkt *probing.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
	}
	pinger.OnFinish = func(stats *probing.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %d duplicates, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	pinger.Count = VarPingCountInt
	pinger.Interval = time.Duration(VarPingInterval) * time.Second
	pinger.Timeout = time.Duration(VarPingTimeout) * time.Second
	pinger.SetPrivileged(VarPingPrivileged)

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		fmt.Println("Failed to ping target host:", err)
	}
	return nil
}
