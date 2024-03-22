package server

import (
	"apus-relayer/relayer/config"
	"apus-relayer/relayer/server/routers"
	"fmt"
)

func Root() {
	r := routers.SetupRouter()
	host := fmt.Sprintf("0.0.0.0:%d", config.GetInt("server.port"))
	if err := r.Run(host); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
