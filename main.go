package main

import (
	"fmt"
	"os"

	scanner "vidar-scan/Scanner" // 导入你的 getscan 包
)

func main() {
	// os.Args 是一个包含所有命令行参数的列表
	// os.Args[0] 是程序的名称（例如 "main" 或 "vidar-scan"）
	// os.Args[1] 是探测对象(端口/路径)

	// os.Args[2] 是我们想要的 URL
	// os.Args[2] 是我们想要的文件名(可选)

	if len(os.Args) < 2 {
		fmt.Println("错误: 未提供参数！")
		fmt.Println("使用方法: go run . <Target-Resource> <Target-Url> <Dictionary-File>")
		fmt.Println("例如: go run . http://example.com/ /path/to/dict.txt")
		os.Exit(1) // 退出程序
	}

	switch os.Args[1] {
	case "PortScan":
		if len(os.Args) != 3 { // check the number of arguments
			fmt.Println("错误: 参数不足！")
			fmt.Println("使用方法: go run . PortScan <target-url>")
			fmt.Println("例如: go run . PortScan http://example.com/")
			os.Exit(1) // 退出程序
		}

		targetUrl := os.Args[2]

		fmt.Printf("[INFO] 开始端口扫描...\n")
		fmt.Printf("[INFO] 目标 URL: %s\n", targetUrl)

		scanner.PortScan(targetUrl)

		fmt.Printf("[INFO] 端口扫描结束。\n")

	case "DirScan":
		if len(os.Args) != 4 { // check the number of arguments
			fmt.Println("错误: 参数不足！")
			fmt.Println("使用方法: go run . DirScan <target-url> <dictionary-file>")
			fmt.Println("例如: go run . DirScan http://example.com/ /path/to/dict.txt")
			os.Exit(1) // 退出程序
		}

		targetUrl := os.Args[2]
		dictFilename := os.Args[3]

		fmt.Printf("[INFO] 开始目录扫描...\n")
		fmt.Printf("[INFO] 目标 URL: %s\n", targetUrl)
		fmt.Printf("[INFO] 使用字典: %s\n", dictFilename)

		scanner.Getscan(targetUrl, dictFilename)

		fmt.Printf("[INFO] 目录扫描结束。\n")

	}
}
