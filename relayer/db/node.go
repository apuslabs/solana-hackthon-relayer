package db

import (
	"apus-relayer/relayer/model"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

const NODE_LOAD = "node_load"
const NODE_INFO = "node_info"

// 合约拉取节点信息，添加node信息到hmap中   考虑是否有剔除节点需要主动通知relayer
func AddNode(node model.GpuNode) error {
	ctx := context.Background()
	_, err := rdb.HSet(ctx, NODE_INFO, node.Id, &node).Result()
	if err != nil {
		fmt.Printf("AddNode: 添加节点失败, 节点信息: %v, 错误信息: %s", node, err.Error())
		return err
	}
	return nil
}

// 0 健康, 1负载过高 2离线(healthcheck = err)
func UpdateNodeLoad(id string, health int) {
	ctx := context.Background()
	_, err := rdb.ZAdd(ctx, NODE_LOAD, &redis.Z{Member: id, Score: float64(health)}).Result()
	if err != nil {
		fmt.Printf("AddNodeHealth: 更新%s Health数据失败, 失败原因: %s\n", id, err.Error())
	}
}

func GetHealthNode() (model.GpuNode, error) {
	ctx := context.Background()
	op := redis.ZRangeBy{
		Min:    "0",
		Max:    "0",
		Offset: 0,
		Count:  1,
	}
	result, err := rdb.ZRangeByScore(ctx, NODE_LOAD, &op).Result()
	if err != nil {
		fmt.Printf("GetHealthNode: 获取健康节点失败, 失败原因: %s\n", err.Error())
		return model.GpuNode{}, err
	}
	if len(result) == 0 {
		fmt.Printf("GetHealthNode: 无健康节点\n")
		return model.GpuNode{}, errors.New("无健康节点")
	}
	nodeId := result[0]

}

func GetNode(id string) model.GpuNode {
	ctx := context.Background()
	result, err := rdb.HGet(ctx, NODE_INFO, id).Result()
}

// 任务如何搞
func AddMetrics(pubkey string, task model.AiTask) {

}

func GetMetrics(pubkey string) []model.AiTask {

}
