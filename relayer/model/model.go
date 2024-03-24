package model

import (
	"encoding/json"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UserInfo struct {
	PubKey string `json:"pubkey"`
	Ctime  int64  `json:"ctime"`
}

func (u *UserInfo) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *UserInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

type Ranking struct {
	Scope  int64  `json:"scope"`
	Pubkey string `json:"pubkey"`
}

type CliHealthResponse struct {
	Port int64 `json:"port"`
	Busy bool  `json:"busy"`
}

type NodeInfo struct {
	Id       string `json:"id"`
	Owner    string `json:"owner"`
	Price    int64  `json:"price"`
	Endpoint string `json:"endpoint"`
}

type GpuNode struct {
	Id          string `json:"id"`
	Owner       string `json:"owner"`
	Cards       []Card `json:"cards"`
	CudaVersion string `json:"cuda_version"`
	Price       int64  `json:"price"`
	Endpoint    string `json:"endpoint"`
}

type Card struct {
	Name   string `json:"name"`
	Memory int    `json:"memory"`
}

func (u *GpuNode) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *GpuNode) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

type AiTask struct {
	User       string              `json:"user"`
	AgentOwner string              `json:"agent_owner"`
	GpuOwner   string              `json:"gpu_owner"`
	Timestamp  timestamp.Timestamp `json:"timestamp"`
	Price      int                 `json:"price"`
}
