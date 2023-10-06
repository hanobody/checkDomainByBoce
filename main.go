package main

import (
	"CHECKDO/myapi"
	"CHECKDO/tools"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

func main() {
	tools.InitConfig()

	for {
		token := "889fd8a2da614502f3b7d623c1c1bb81"

		// 从文本文件中读取域名列表
		filename := "hosts.txt"
		hosts, err := tools.ReadHostsFromFile(filename)
		if err != nil {
			fmt.Printf("无法读取域名列表文件：%v\n", err)
			return
		}

		// 调用 CreateTasks 函数来创建多个任务
		taskID := tools.CreateTasks(token, hosts)

		// 沉睡1分钟等待结果
		var waitreport = viper.GetInt("timeconfig.waitreport")
		time.Sleep(time.Minute * time.Duration(waitreport))

		// 获取 JSON 数据
		responses, err := myapi.FetchJSONData(taskID, token)
		if err != nil {
			fmt.Println("获取检测结果的 JSON 数据时出错:", err)
			return
		}

		// 创建或打开文本文件
		file, err := os.Create("results.txt")
		if err != nil {
			fmt.Println("无法创建或打开结果文件:", err)
			return
		}
		defer file.Close()

		// 遍历每个响应
		for i, response := range responses {
			output := "\n" + response.List[i].Host + "\n"
			for _, listItem := range response.List {
				if listItem.HTTPCode == 0 {
					output += fmt.Sprintf("状态: 000,%s\n", listItem.NodeName)
				} else {
					if listItem.RemoteIP == "127.0.0.1" {
						output += fmt.Sprintf("%s,%s\n", listItem.RemoteIP, listItem.NodeName)
					} else if tools.ContainsIP(tools.CheckSafe("antispam.txt"), listItem.RemoteIP) {
						output += fmt.Sprintf("反诈中心:%s,%s\n", listItem.NodeName, listItem.RemoteIP)
					} else if tools.ContainsIP(tools.CheckSafe("hijack.txt"), listItem.RemoteIP) {
						output += fmt.Sprintf("被劫持了:%s,%s\n", listItem.NodeName, listItem.RemoteIP)
					}
				}

			}
			// 将输出写入文件
			tools.TelegramSender(output)
			fmt.Fprint(file, output)
		}
		// 沉睡2小时等下次检测
		var nextaction = viper.GetInt("time.nextaction")
		time.Sleep(time.Minute * time.Duration(nextaction))
	}
}
