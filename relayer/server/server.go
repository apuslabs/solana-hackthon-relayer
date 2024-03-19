package server

import "fmt"

func Root() {
	r := setupRouter()
	if err := r.Run("0.0.0.0:80"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
