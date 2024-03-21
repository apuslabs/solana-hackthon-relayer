package server

import (
	"apus-relayer/relayer/server/routers"
	"fmt"
)

func Root() {
	r := routers.SetupRouter()
	if err := r.Run("0.0.0.0:80"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
