package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// 定义一个结构来匹配JSON响应的格式
type GetNodes struct {
	ErrorCode int    `json:"error_code"`
	Error     string `json:"error"`
	Data      struct {
		List []Node `json:"list"`
	} `json:"data"`
}

type Node struct {
	ID       int    `json:"id"`
	NodeName string `json:"node_name"`
	ISPName  string `json:"isp_name"`
	ISPCode  int    `json:"isp_code"`
}

func ProcessNodes(token string) string {
	// 定义要请求的URL
	url := "https://task.boce.com/v3/node/list?key=" + token

	// 发送HTTP GET请求
	//response, err := http.Get(url)
	response, err := http.Post(url, "application/json", nil)
	if err != nil {
		return "HTTP请求失败: " + err.Error()
	}
	defer response.Body.Close()

	// 解码JSON响应
	var jsonResponse GetNodes
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&jsonResponse); err != nil {
		return "无法解码JSON响应: " + err.Error()
	}

	// 检查是否存在错误
	if jsonResponse.ErrorCode != 0 {
		return fmt.Sprintf("错误代码: %d\n错误信息: %s", jsonResponse.ErrorCode, jsonResponse.Error)
	}

	// 构建拼接后的ID字符串
	var idStrings []string
	var nodesNumber int
	for _, node := range jsonResponse.Data.List {
		if node.ISPName != "海外" && node.ISPName != "港澳台" {
			idStrings = append(idStrings, fmt.Sprintf("%d", node.ID))
			nodesNumber++
		}
	}
	fmt.Println("当前境内节点数量为:", nodesNumber)
	// 使用逗号分隔符将ID字符串连接起来
	idConcatenated := strings.Join(idStrings, ",")

	return idConcatenated
}
