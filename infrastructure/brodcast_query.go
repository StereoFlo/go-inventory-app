package infrastructure

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type IP struct {
	IP        string
	Broadcast string
}

func SendBroadcast(sleep time.Duration) error {
	packetConn, err := net.ListenPacket("udp4", ":2712")
	if err != nil {
		return err
	}
	defer packetConn.Close()

	for {
		ips, err := getIpData()
		if err != nil {
			return err
		}

		for _, v := range ips {
			addr, err := net.ResolveUDPAddr("udp4", v.Broadcast+":2712")
			if err != nil {
				return err
			}
			_, err = packetConn.WriteTo([]byte(v.IP+":"+os.Getenv("API_PORT")), addr)
			if err != nil {
				return err
			}
		}
		time.Sleep(sleep)
	}
}

func getIpBroadcast(subnet *net.IPNet) net.IP {
	ipLen := len(subnet.IP)
	out := make(net.IP, ipLen)
	var m byte
	for i := 0; i < ipLen; i++ {
		m = subnet.Mask[i] ^ 0xff
		out[i] = subnet.IP[i] | m

	}
	return out
}

func getIpData() ([]*IP, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var ips []*IP
	for _, i := range interfaces {
		ips, err = getIps(i)
		if err != nil {
			log.Println(fmt.Errorf("localAddresses: %+v\n", err.Error()))
			return nil, err
		}
	}

	return ips, nil
}

func getIps(i net.Interface) ([]*IP, error) {
	addrs, err := i.Addrs()
	if err != nil {
		return nil, err
	}
	var ips []*IP
	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPNet:
			inet, ok := addr.(*net.IPNet)
			if ok && !v.IP.IsLoopback() && inet.IP.To4() != nil {
				_, ipnet, _ := net.ParseCIDR(v.String())
				broadcast := getIpBroadcast(ipnet)
				ip := &IP{
					IP:        inet.IP.To4().String(),
					Broadcast: broadcast.String(),
				}
				ips = append(ips, ip)
			}
		}

	}
	return ips, err
}
