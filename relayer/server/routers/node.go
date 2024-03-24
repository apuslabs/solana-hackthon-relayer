package routers

import (
	"apus-relayer/relayer/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func initNode(r *gin.Engine) {
	r.POST("/registerNode", RegisterNodeHandler)
	r.POST("/matchNode", MatchNodeHandler)
	r.POST("/submitMetrics", SubmitMetricsHandler)
}

// 节点cli启动后，调用该接口注册cli信息
func RegisterNodeHandler(c *gin.Context) {
	var gpuNode model.GpuNode
	err := c.ShouldBindBodyWith(&gpuNode, binding.JSON)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{Code: 500, Msg: err.Error(), Data: ""})
		return
	}
	// 上传节点信息
}

// 根据任务，匹配适用节点
func MatchNodeHandler(c *gin.Context) {
	// 获取任务信息
	// 查询同一用户是否有其他未确认流水
	// 查询用户是否有消费代币
	// 获取当前健康节点中最便宜的返回
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: "hello"})
}

// 获取用户未上链流水
func GetMetricsHandler(c *gin.Context) {

}

// 提交用户流水，缓存流水
func SubmitMetricsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: "hello"})
}
