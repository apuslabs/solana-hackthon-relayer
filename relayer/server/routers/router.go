package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	loadRouters(r)
	return r
}

func loadRouters(r *gin.Engine) {
	initReferral(r)
	initUser(r)
	initNode(r)
}
