package basework

import (
	"strings"
)

// 参数处理也全部丢到这个文件里面吧

func ParseArgs(s []string) (map[string]string, error) {
	ArgsMap := make(map[string]string)

	var Args = ""
	var exist = false
	for _, value := range s {
		if exist {
			ArgsMap[Args] = value
			exist = false
		}

		if strings.HasPrefix(value, "-") {
			//fmt.Println(i, value)
			Args = value[1:]
			exist = true
		}

	}

	return ArgsMap, nil
}

func DealPort(s string) (int, int, error) {
	BeginPort := 0
	EndPort := 65535

	return BeginPort, EndPort, nil
}
