package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 定义一个结构来匹配JSON响应的格式
//
//	type CreateTask struct {
//		ErrorCode int    `json:"error_code"`
//		Error     string `json:"error"`
//		Data      struct {
//			ID string `json:"id"`
//		} `json:"data"`
//	}
type CreateTask struct {
	ErrorCode int    `json:"error_code"`
	Data      Data   `json:"data"`
	Error     string `json:"error"`
}
type Data struct {
	ID string `json:"id"`
}

// 创建多个任务的函数
func CreateTasks(token string, hosts []string) []string {
	var taskIDs []string // 用于存储每个域名的任务ID

	if len(hosts) == 0 {
		fmt.Println("域名列表为空")
		return taskIDs // 返回空切片
	}

	// 获取节点信息
	nodes := ProcessNodes(token)

	// 检查节点信息是否为空
	if nodes == "" {
		fmt.Println("节点信息为空")
		return taskIDs // 返回空切片
	}

	// 创建任务并处理每个域名
	for _, host := range hosts {
		// 定义要请求的URL
		url := "https://api.boce.com/v3/task/create/curl?key=" + token + "&node_ids=" + nodes + "&host=" + host

		response, err := http.Post(url, "application/json", nil)
		if err != nil {
			fmt.Printf("处理域名 %s 时发生HTTP请求错误: %v\n", host, err)
			continue
		}
		defer response.Body.Close()

		// 解码JSON响应
		var jsonResponse CreateTask
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&jsonResponse); err != nil {
			fmt.Printf("处理域名 %s 时无法解码JSON响应: %v\n", host, err)
			continue
		}

		// 检查是否存在错误
		if jsonResponse.ErrorCode != 0 {
			fmt.Printf("处理域名 %s 时发生错误: 错误代码: %d, 错误信息: %s\n", host, jsonResponse.ErrorCode, jsonResponse.Error)
			continue
		}

		//fmt.Println("创建的任务ID:" + jsonResponse.Data.ID)
		// 将任务ID添加到切片中
		taskIDs = append(taskIDs, jsonResponse.Data.ID)
	}

	return taskIDs
}
