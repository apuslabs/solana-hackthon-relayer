package server

import (
	"apus-relayer/relayer/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	loadRouters(r)
	return r
}

func loadRouters(r *gin.Engine) {
	r.POST("/getNode", getNodeHandler)
	r.POST("/submitMetrics", submitMetricsHandler)
}

func getNodeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: "hello"})
}

func submitMetricsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: "hello"})
}