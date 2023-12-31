package http

import (
	"fmt"
	"github.com/kataras/iris/v12/core/netutil"
	"homeHelper/config"
	"net"
)

func FindPublicAddresses() ([]string, error) {
	ips, err := findLocalIps()
	var result []string

	if err != nil {
		return result, err
	}

	for _, ip := range ips {
		//listeningURI := netutil.ResolveURL("http", fmt.Sprintf("%s:%s", ip, config.APP["APP_PORT"]))
		result = append(result, netutil.ResolveURL("http", fmt.Sprintf("%s:%s", ip, config.APP["APP_PORT"])))
	}

	return result, nil
}

func findLocalIps() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var ips []string
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, localAddr := range addrs {
			var ip net.IP
			switch v := localAddr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip != nil && ip.To4() != nil {
				if !ip.IsPrivate() {
					continue
				}
				ips = append(ips, ip.String())
			}
		}
	}

	return ips, nil
}
