package basework

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// 参数处理全部丢到这个文件里面吧

func ParsePort(s string) (int, int, error) {
	StartPort, EndPort := 0, 65535
	var err error

	PartsPort := strings.Split(s, "-")
	//fmt.Println(PartsPort)

	if len(PartsPort) == 1 {
		if port, err := strconv.Atoi(PartsPort[0]); err == nil {
			return port, port, nil
		}
		return 0, 0, fmt.Errorf("Parse error: invalid port number")
	}

	if len(PartsPort) != 2 {
		return 0, 0, fmt.Errorf("Parse error: invalid port number")
	}

	StartPort, err = strconv.Atoi(PartsPort[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Parse error: invalid start port")
	}

	EndPort, err = strconv.Atoi(PartsPort[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Parse error: invalid end port")
	}

	if StartPort < 0 || StartPort > EndPort || EndPort > 65535 {
		return 0, 0, fmt.Errorf("Parse error: invalid port number")
	}
	return StartPort, EndPort, nil
}

func ParseCIDR(cidr string) (string, string, error) {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", "", fmt.Errorf("invalid CIDR: %v", err)
	}

	StartIP := ipnet.IP.Mask(ipnet.Mask)

	ip := ipnet.IP.To4()
	mask := ipnet.Mask

	broadcast := make(net.IP, len(ip))
	for i := 0; i < 4; i++ {
		broadcast[i] = ip[i] | ^mask[i]
	}

	EndIP := broadcast

	return StartIP.String(), EndIP.String(), nil
}
