package tools

import (
	"fmt"
	"net"
	"strings"
)

func CheckSafe(ipList string) []string {
	// 从文本文件中读取域名列表
	ips, err := ReadHostsFromFile(ipList)
	if err != nil {
		fmt.Printf("无法读取域名列表文件：%v\n", err)
		return nil
	}
	return ips
}

// containsIP 函数用于检查字符串是否存在于包含具体 IP 地址和 IP 段的切片中
func ContainsIP(slice []string, target string) bool {
	for _, str := range slice {
		// 使用 strings.Contains 检查 target 是否在 str 中
		if strings.Contains(str, target) {
			return true
		}

		// 尝试解析为 IP 地址或 IP 段
		_, ipNet, err := net.ParseCIDR(str)
		if err != nil {
			// 解析失败，可能是具体的 IP 地址
			ip := net.ParseIP(str)
			if ip != nil && ip.String() == target {
				return true
			}
		} else {
			// 解析成功，是一个 IP 段
			if ipNet.Contains(net.ParseIP(target)) {
				return true
			}
		}
	}
	return false
}
