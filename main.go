package main

import (
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
