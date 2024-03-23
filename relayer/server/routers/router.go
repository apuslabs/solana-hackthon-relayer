package routers

import (
	"apus-relayer/relayer/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	loadRouters(r)
	return r
}

func loadRouters(r *gin.Engine) {
	initReferral(r)
	initUser(r)
	// node routers
	r.POST("/getNode", getNodeHandler)
	r.POST("/submitMetrics", submitMetricsHandler)

	// referral routers
}

func getNodeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: "hello"})
}

func submitMetricsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: "hello"})
}
