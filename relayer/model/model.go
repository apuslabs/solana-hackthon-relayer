package model

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

type NodeInfo struct {
}
