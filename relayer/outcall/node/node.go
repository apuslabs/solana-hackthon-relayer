package node

import (
	"apus-relayer/relayer/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HealthCheck(host string, port int, hash string) model.CliHealthResponse {
	url := fmt.Sprintf("http://%s:%d/status?hash=%s", host, port, hash)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("HealthCheck: 获取节点%s健康状态失败, 失败原因: %s", host, err.Error())
		return model.CliHealthResponse{Port: 0, Busy: true}
	}
	defer resp.Body.Close()
	bytes, _ := io.ReadAll(resp.Body)
	var result model.CliHealthResponse
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return model.CliHealthResponse{Port: 0, Busy: true}
	}
	return result
}
