package model

import "encoding/json"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UserInfo struct {
	PubKey string `json:"pubkey"`
	Code   string `json:"code"`
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

type NodeInfo struct {
}
