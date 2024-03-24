package routers

import (
	"apus-relayer/relayer/db"
	"apus-relayer/relayer/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initUser(r *gin.Engine) {
	r.POST("/user/signup", SignupHandler)
}

// 用户登记 记录第一次使用时间 绑定邀请关系
func SignupHandler(c *gin.Context) {

	pubkey := c.PostForm("publickey")
	referralkey := c.PostForm("referralkey")

	if pubkey == "" {
		c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: ""})
		return
	}

	user, err := db.GetUser(pubkey)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{Code: 500, Msg: err.Error(), Data: ""})
		return
	}

	if user.PubKey != "" {
		c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: user})
	}

	// 用户未注册,注册用户信息
	user.PubKey = pubkey
	user.Ctime = time.Now().Unix()
	if err := db.AddUser(user); err != nil {
		c.JSON(http.StatusOK, model.Response{Code: 500, Msg: err.Error(), Data: ""})
		return
	}
	if referralkey != "" {
		db.AddReferral(referralkey, user)
	}
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: user})
}
