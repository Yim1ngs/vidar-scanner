package main

import (
	"fmt"
	"os"
	"vidar-scan/Scanner"
	"vidar-scan/basework"
)

func main() {
	/*
		os.Args 是一个包含所有命令行参数的列表
		os.Args[0] 是程序的名称（例如 "main" 或 "vidar-scan"）
		os.Args[1] 是探测对象(端口/路径)
		os.Args[2] 是我们想要的 URL
		os.Args[3] 是我们想要的文件名(可选)
	*/

	if len(os.Args) < 2 {
		fmt.Println("错误: 未提供参数！")
		fmt.Println("使用方法: go run . <Target-Resource>")
		fmt.Println("例如: go run . PortScan/DirScan")
		os.Exit(1) // 退出程序
	}

	ArgsMap, err := basework.ParseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch os.Args[1] {

	case "PortScan":
		NeedArgs := []string{"u"}
		for _, arg := range NeedArgs {
			_, exists := ArgsMap[arg]
			if !exists {
				fmt.Println("Missing Args: u")
				fmt.Println("使用方法: go run . PortScan -u <target-url>")
				fmt.Println("例如: go run . PortScan -u http://example.com/")
				os.Exit(1)
			}
		}

		targetUrl := ArgsMap["u"]
		var BeginPort, EndPort = 0, 65535

		_, exists := ArgsMap["p"]
		if exists {
			BeginPort, EndPort, err = basework.DealPort(ArgsMap["p"])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		fmt.Printf("[INFO] 开始端口扫描...\n")
		fmt.Printf("[INFO] 目标 URL: %s\n", targetUrl)
		fmt.Printf("[INFO] 端口范围: %d-%d\n", BeginPort, EndPort)

		scanner.PortScan(targetUrl, BeginPort, EndPort)

		fmt.Printf("[INFO] 端口扫描结束。\n")

	case "DirScan":
		NeedArgs := []string{"u", "d"}
		for _, arg := range NeedArgs {
			_, exists := ArgsMap[arg]
			if !exists {
				fmt.Printf("Missing Args: %s\n", arg)
				fmt.Println("使用方法: go run . PortScan -u <target-url> -d <dictionary>")
				fmt.Println("例如: go run . PortScan -u http://example.com/ -d /path/to/dict")
				os.Exit(1)
			}
		}

		targetUrl, _ := ArgsMap["u"]
		dictFilename, _ := ArgsMap["d"]

		fmt.Printf("[INFO] 开始目录扫描...\n")
		fmt.Printf("[INFO] 目标 URL: %s\n", targetUrl)
		fmt.Printf("[INFO] 使用字典: %s\n", dictFilename)

		scanner.Getscan(targetUrl, dictFilename)

		fmt.Printf("[INFO] 目录扫描结束。\n")

	default:
		fmt.Printf("[ERROR] 未定义的探测对象")
	}
}
