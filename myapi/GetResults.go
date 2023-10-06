package myapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Done    bool       `json:"done"`
	ID      string     `json:"id"`
	List    []ListItem `json:"list"`
	MaxNode int        `json:"max_node"`
}

type ListItem struct {
	NodeID         int     `json:"node_id"`
	NodeName       string  `json:"node_name"`
	Host           string  `json:"host"`
	OriginIP       string  `json:"origin_ip"`
	RemoteIP       string  `json:"remote_ip"`
	IPRegion       string  `json:"ip_region"`
	HTTPCode       int     `json:"http_code"`
	TimeTotal      float64 `json:"time_total"`
	TimeNamelookup float64 `json:"time_namelookup"`
	TimeConnect    float64 `json:"time_connect"`
	SizeDownload   float64 `json:"size_download"`
	SpeedDownload  float64 `json:"speed_download"`
	DownloadTime   float64 `json:"download_time"`
	ReportSource   string  `json:"report_source"`
	ErrorCode      int     `json:"error_code"`
	Error          string  `json:"error"`
}

func FetchJSONData(taskIDs []string, token string) ([]*Result, error) {
	var responses []*Result

	for _, taskID := range taskIDs {
		url := fmt.Sprintf("https://api.boce.com/v3/task/curl/%s?key=%s", taskID, token)

		resp, err := http.Post(url, "application/json", nil)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// 读取响应体
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// 解析JSON数据
		var response Result
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}

		responses = append(responses, &response)
	}

	return responses, nil
}
