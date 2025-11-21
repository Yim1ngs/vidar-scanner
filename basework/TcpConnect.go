package basework

import (
	"fmt"
	"net"
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
