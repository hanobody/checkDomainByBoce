package tools

import (
	"bufio"
	"os"
	"strings"
)

// 从文本文件中读取域名列表
func ReadHostsFromFile(filename string) ([]string, error) {
	var ips []string

	file, err := os.Open(filename)
	if err != nil {
		return ips, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		host := strings.TrimSpace(scanner.Text())
		if host != "" {
			ips = append(ips, host)
		}
	}

	if err := scanner.Err(); err != nil {
		return ips, err
	}

	return ips, nil
}
