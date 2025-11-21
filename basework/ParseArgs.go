package basework

import (
	"fmt"
	"strconv"
	"strings"
)

// 参数处理全部丢到这个文件里面吧

func DealPort(s string) (int, int, error) {
	BeginPort := 0
	EndPort := 65535

	PartsPort := strings.Split(s, "-")
	//fmt.Println(PartsPort)

	if len(PartsPort) != 2 {
		return 0, 0, fmt.Errorf("Parse error: invalid port number")
	}

	BeginPort, _ = strconv.Atoi(PartsPort[0])
	EndPort, _ = strconv.Atoi(PartsPort[1])

	if BeginPort < 0 || BeginPort > EndPort || EndPort > 65535 {
		return 0, 0, fmt.Errorf("Parse error: invalid port number")
	}
	return BeginPort, EndPort, nil
}

func DealCIDRIP(s string) (string, int, error) {
	parts := strings.Split(s, "/")
	ip := parts[0]
	CIDR, err := strconv.Atoi(parts[1])

	if err != nil {
		return "", 0, err
	}

	fmt.Println(CIDR)

	if !(CIDR%8 == 0 && CIDR <= 24 && CIDR >= 0) {
		return "", 0, fmt.Errorf("Parse error: invalid CIDR IP")
	}

	return ip, CIDR, nil
}
