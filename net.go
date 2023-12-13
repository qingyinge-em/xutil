package xutil

import (
	"net"
	"time"

	"github.com/spf13/cast"
)

// param: 1.2.3.4:443
func IsIpPortValid1(ipport string) bool {
	ip, port, err := net.SplitHostPort(ipport)
	if err != nil {
		return false
	}
	return IsIpPortValid2(ip, cast.ToInt(port))
}

func IsIpPortValid2(ip string, port int) bool {
	if net.ParseIP(ip) == nil {
		return false
	}
	p := cast.ToInt(port)
	if p < 1 || p > 65535 {
		return false
	}
	return true
}

// 端口探测
func PortTest(ip string, port int, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", ip+":"+cast.ToString(port), timeout)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}
