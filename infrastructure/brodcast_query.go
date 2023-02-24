package infrastructure

import (
	"net"
	"os"
	"strconv"
	"time"
)

type ip struct {
	IP        string
	Broadcast string
}

func SendBroadcast() error {
	sleep, err := strconv.Atoi(os.Getenv("BROADCAST_SLEEP"))
	if err != nil {
		return err
	}
	packetConn, err := net.ListenPacket("udp4", ":"+os.Getenv("BROADCAST_PORT"))
	if err != nil {
		return err
	}
	defer packetConn.Close()

	for {
		ips, err := getInterfacesIPs()
		if err != nil {
			return err
		}

		for _, v := range ips {
			addr, err := net.ResolveUDPAddr("udp4", v.Broadcast+":"+os.Getenv("BROADCAST_PORT"))
			if err != nil {
				return err
			}
			_, err = packetConn.WriteTo([]byte(v.IP+":"+os.Getenv("API_PORT")), addr)
			if err != nil {
				return err
			}
		}
		time.Sleep(time.Second * time.Duration(sleep))
	}
}

func getBroadcastIp(subnet *net.IPNet) net.IP {
	ipLen := len(subnet.IP)
	out := make(net.IP, ipLen)
	var m byte
	for i := 0; i < ipLen; i++ {
		m = subnet.Mask[i] ^ 0xff
		out[i] = subnet.IP[i] | m

	}
	return out
}

func getInterfacesIPs() ([]*ip, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var ips []*ip
	for _, i := range interfaces {
		ips, err = getLocalIPs(i)
		if err != nil {
			return nil, err
		}
	}

	return ips, nil
}

func getLocalIPs(i net.Interface) ([]*ip, error) {
	addrs, err := i.Addrs()
	if err != nil {
		return nil, err
	}
	var ips []*ip
	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPNet:
			inet, ok := addr.(*net.IPNet)
			if ok && !v.IP.IsLoopback() && inet.IP.To4() != nil {
				_, ipnet, _ := net.ParseCIDR(v.String())
				broadcast := getBroadcastIp(ipnet)
				ip := &ip{
					IP:        inet.IP.To4().String(),
					Broadcast: broadcast.String(),
				}
				ips = append(ips, ip)
			}
		}

	}
	return ips, err
}
