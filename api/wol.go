package api

import (
	"errors"
	"fmt"
	"github.com/go-ping/ping"
	"github.com/sabhiram/go-wol/wol"
	"net"
	"time"
)

// Wol 网络唤醒
var Wol = new(_Wol)

type _Wol struct {
}

//WakeOnLan 唤醒局域网内的电脑
func (c *_Wol) WakeOnLan(macAddr, face, ip string, isJumpIpCheck bool) error {
	if !isJumpIpCheck {
		isPowerOn, err := c.IsPowerOn(ip)
		if err != nil {
			return err
		}
		if isPowerOn {
			return errors.New("当前电脑已经处于开机状态了")
		}
	}

	var localAddr *net.UDPAddr
	localAddr, _ = c.ipFromInterface(face)
	//&ping.Statistics{PacketsRecv:0, PacketsSent:1, PacketsRecvDuplicates:0, PacketLoss:100, IPAddr:(*net.IPAddr)(0xc0003bd1d0), Addr:"192.168.3.77", Rtts:[]time.Duration(nil), MinRtt:0, MaxRtt:0, AvgRtt:0, StdDevRtt:0}
	//&ping.Statistics{PacketsRecv:1, PacketsSent:1, PacketsRecvDuplicates:0, PacketLoss:0, IPAddr:(*net.IPAddr)(0xc000145470), Addr:"192.168.3.81", Rtts:[]time.Duration{3447000}, MinRtt:3447000, MaxRtt:3447000, AvgRtt:3447000, StdDevRtt:0}
	udpAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")
	if err != nil {
		return err
	}
	mp, err := wol.New(macAddr)
	if err != nil {
		return err
	}

	bs, err := mp.Marshal()
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", localAddr, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	n, err := conn.Write(bs)
	if err == nil && n != 102 {
		return errors.New(fmt.Sprintf("发送魔术封包失败,状态码: %v", n))
	}
	return nil
}

//IsPowerOn 判断目标机器是否开机
func (c *_Wol) IsPowerOn(ip string) (bool, error) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return false, err
	}
	pinger.Count = 1
	pinger.Timeout = time.Second
	if err := pinger.Run(); err != nil {
		//如果出现没有权限需要运行 sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"
		return false, errors.New(fmt.Sprintf("检测网络失败: %s", err.Error()))
	}
	stats := pinger.Statistics()
	if stats.PacketsRecv >= 1 {
		return true, nil
	}
	return false, nil //没开机
}
func (_Wol) ipFromInterface(iface string) (*net.UDPAddr, error) {
	ief, err := net.InterfaceByName(iface)
	if err != nil {
		return nil, err
	}

	addrs, err := ief.Addrs()
	if err == nil && len(addrs) <= 0 {
		err = fmt.Errorf("没有获取到当前网卡IP地址: %s", iface)
	}
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		switch ip := addr.(type) {
		case *net.IPNet:
			if ip.IP.DefaultMask() != nil {
				return &net.UDPAddr{
					IP: ip.IP,
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("没有找到到当前网卡IP地址: %s", iface)
}
