/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	scanner "vidar-scan/Scanner"
	"vidar-scan/basework"

	"github.com/spf13/cobra"
)

var (
	IpHost string
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "host scanning",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		StartIP, EndIP, err := basework.ParseCIDR(IpHost)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//fmt.Println(StartIP, EndIP)

		fmt.Printf("[INFO] 开始Host扫描...\n")
		fmt.Printf("[INFO] Host范围: %s-%s\n", StartIP, EndIP)

		scanner.HostScan(StartIP, EndIP)
		fmt.Printf("[INFO] Host扫描结束。\n")
	},
}

func init() {
	rootCmd.AddCommand(hostCmd)

	hostCmd.Flags().StringVarP(&IpHost, "ip", "i", "", "Target IPv4 (required)")

	hostCmd.MarkFlagRequired("ip")
}
