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
	pc, err := net.ListenPacket("udp4", ":2712")
	if err != nil {
		return err
	}
	defer pc.Close()

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
			_, err = pc.WriteTo([]byte(v.IP+":"+os.Getenv("API_PORT")), addr)
			if err != nil {
				return err
			}
		}
		time.Sleep(sleep)
	}
}

func getIpBroadcast(subnet *net.IPNet) net.IP {
	n := len(subnet.IP)
	out := make(net.IP, n)
	var m byte
	for i := 0; i < n; i++ {
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
		addrs, err := i.Addrs()
		if err != nil {
			log.Println(fmt.Errorf("localAddresses: %+v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPNet:
				inet, ok := a.(*net.IPNet)
				if ok && !v.IP.IsLoopback() && inet.IP.To4() != nil {
					_, ipnet, _ := net.ParseCIDR(v.String())
					br := getIpBroadcast(ipnet)
					ip := &IP{
						IP:        inet.IP.To4().String(),
						Broadcast: br.String(),
					}
					ips = append(ips, ip)
				}
			}

		}
	}

	return ips, nil
}
