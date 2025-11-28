package basework

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func TcpConnect(ip string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err == nil {
		//fmt.Println(port)
		conn.Close()
		return true
	} else {
		//fmt.Printf("error type: %T, message: %v\n", err, err)
	}
	return false
}

func IsAliveTCP(ip string, timeout time.Duration) bool {
	done := make(chan bool)

	for _, port := range []int{80, 443, 8080} {

		go func(p int) {

			addr := fmt.Sprintf("%s:%d", ip, p)
			conn, err := net.DialTimeout("tcp", addr, timeout)

			if err == nil {
				conn.Close()
				done <- true
				return
			}

			if err != nil && strings.Contains(err.Error(), "refused") {
				done <- true
				return
			}

		}(port)
	}

	select {
	case <-done:
		return true
	case <-time.After(timeout + 100*time.Millisecond): // 稍微多等一点点
		return false

	}
}
