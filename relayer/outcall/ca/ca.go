package ca

import (
	"apus-relayer/relayer/model"
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
	"log"
)

const PUB_KEY = "3jDFmwPSPScHvGhzLwniwd6id232pcJKumszMHWHLEHF"

func main() {
	c := client.NewClient(rpc.TestnetRPCEndpoint)

	// 获取余额
	balance, err := c.GetBalance(context.TODO(), PUB_KEY)
	if err != nil {
		log.Fatalf("failed to get balance, err: %v", err)
	}
	fmt.Printf("balance: %v SOL\n", balance)

}

func RegistNode([]model.GpuNode) bool {
	return false
}

func GetBalance(pubkey string) int64 {
	return 0
}

func GetNodes() []model.GpuNode {
	return []model.GpuNode{}
}

func SubmitMetrics([]model.AiTask) bool {
	return false
}
