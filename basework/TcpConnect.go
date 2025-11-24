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
		// 	    fmt.Println(port)
		conn.Close()
		return true
	} else {
		// 	    fmt.Printf("error type: %T, message: %v\n", err, err)
	}
	return false
}

func IsAliveTCP(ip string, timeout time.Duration) bool {
	for _, port := range []int{80, 443, 8080} {
		addr := fmt.Sprintf("%s:%d", ip, port)
		//fmt.Println(addr)
		conn, err := net.DialTimeout("tcp", addr, timeout)
		//fmt.Println(conn, err)
		if err == nil {
			conn.Close()
			return true
		}

		if strings.Contains(err.Error(), "refused") { //
			return true
		}
	}
	return false
}
